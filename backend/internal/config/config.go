package config

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Env               string
	Port              string
	DBDSN             string
	JWTSecret         string
	AccessTTL         time.Duration
	RefreshTTL        time.Duration
	CORSOrigins       []string
	RunnerMode        string
	PlaygroundURL     string
	DockerImage       string
	RunTimeout        time.Duration
	MaxCodeBytes      int
	MaxOutputBytes    int
	PublicRatePerMin  int
	PlaygroundRateMin int
}

func Load() Config {
	_ = godotenv.Load()

	cfg := Config{
		Env:               env("APP_ENV", "development"),
		Port:              env("PORT", "8080"),
		DBDSN:             env("DB_DSN", "root:torikal@tcp(127.0.0.1:3306)/golearn?parseTime=true&multiStatements=true&charset=utf8mb4"),
		JWTSecret:         env("JWT_SECRET", "dev-secret-change-me-in-production"),
		AccessTTL:         durationEnv("ACCESS_TTL", 15*time.Minute),
		RefreshTTL:        durationEnv("REFRESH_TTL", 7*24*time.Hour),
		CORSOrigins:       splitEnv("CORS_ORIGINS", "http://localhost:5173"),
		RunnerMode:        env("RUNNER_MODE", "playground"),
		PlaygroundURL:     env("PLAYGROUND_URL", "https://go.dev/_/compile"),
		DockerImage:       env("RUNNER_DOCKER_IMAGE", "golang:1.22-alpine"),
		RunTimeout:        durationEnv("RUN_TIMEOUT", 8*time.Second),
		MaxCodeBytes:      intEnv("MAX_CODE_BYTES", 64*1024),
		MaxOutputBytes:    intEnv("MAX_OUTPUT_BYTES", 32*1024),
		PublicRatePerMin:  intEnv("RATE_LIMIT_PER_MIN", 120),
		PlaygroundRateMin: intEnv("PLAYGROUND_RATE_PER_MIN", 20),
	}
	if cfg.Env == "production" && cfg.JWTSecret == "dev-secret-change-me-in-production" {
		panic("JWT_SECRET must be set in production")
	}
	return cfg
}

func env(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

func durationEnv(key string, fallback time.Duration) time.Duration {
	v, err := time.ParseDuration(env(key, fallback.String()))
	if err != nil {
		return fallback
	}
	return v
}

func intEnv(key string, fallback int) int {
	v, err := strconv.Atoi(env(key, strconv.Itoa(fallback)))
	if err != nil {
		return fallback
	}
	return v
}

func splitEnv(key, fallback string) []string {
	raw := env(key, fallback)
	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}
