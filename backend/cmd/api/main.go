package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golearn/internal/config"
	"golearn/internal/database"
	"golearn/internal/handler"
	"golearn/internal/router"
	"golearn/internal/runner"
	"golearn/internal/seed"
	"golearn/internal/service"
	"golearn/internal/store"
)

func main() {
	cfg := config.Load()
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))

	db, err := database.Open(cfg.DBDSN)
	if err != nil {
		slog.Error("database connection failed", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	if err = database.Migrate(db); err != nil {
		slog.Error("migration failed", "error", err)
		os.Exit(1)
	}
	if err = seed.Run(db); err != nil {
		slog.Error("seed failed", "error", err)
		os.Exit(1)
	}

	st := store.New(db)
	tokens := service.NewTokenService(cfg.JWTSecret, cfg.AccessTTL, cfg.RefreshTTL)
	authSvc := service.NewAuthService(st, tokens)
	contentSvc := service.NewContentService(st)
	meSvc := service.NewMeService(st)

	var codeRunner runner.Runner
	if cfg.RunnerMode == "docker" {
		codeRunner = runner.NewDocker(cfg.DockerImage, cfg.RunTimeout, cfg.MaxCodeBytes, cfg.MaxOutputBytes)
	} else {
		codeRunner = runner.NewPlayground(cfg.PlaygroundURL, cfg.RunTimeout, cfg.MaxCodeBytes, cfg.MaxOutputBytes)
	}
	learnSvc := service.NewLearnService(st, codeRunner)

	h := &handler.Handlers{Auth: authSvc, Content: contentSvc, Learn: learnSvc, Me: meSvc}
	r := router.New(cfg, h, tokens)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		slog.Info("golearn api listening", "port", cfg.Port, "runner", codeRunner.Name(), "env", cfg.Env)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server error", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("shutdown failed", "error", err)
	}
	slog.Info("server stopped")
}
