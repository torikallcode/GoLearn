package router

import (
	"time"

	"github.com/gin-gonic/gin"

	"golearn/internal/config"
	"golearn/internal/handler"
	"golearn/internal/middleware"
	"golearn/internal/response"
	"golearn/internal/service"
)

func New(cfg config.Config, h *handler.Handlers, tokens *service.TokenService) *gin.Engine {
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.CORS(cfg.CORSOrigins))
	r.GET("/healthz", func(c *gin.Context) {
		response.OK(c, gin.H{"status": "ok", "time": time.Now().UTC()})
	})

	api := r.Group("/api/v1")
	api.Use(middleware.RateLimit(cfg.PublicRatePerMin))

	auth := api.Group("/auth")
	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)
	auth.POST("/refresh", h.Refresh)
	auth.POST("/logout", h.Logout)
	auth.GET("/me", middleware.Auth(tokens), h.CurrentUser)
	auth.DELETE("/me", middleware.Auth(tokens), h.DeleteAccount)

	optional := middleware.OptionalAuth(tokens)
	api.GET("/paths", optional, h.Paths)
	api.GET("/paths/:slug", optional, h.Path)
	api.GET("/lessons/:slug", optional, h.Lesson)
	api.GET("/lessons/:slug/exercises", optional, h.LessonExercises)
	api.GET("/lessons/:slug/quizzes", optional, h.LessonQuizzes)
	api.GET("/search", h.Search)
	api.GET("/projects", optional, h.Projects)
	api.GET("/projects/:slug", optional, h.Project)

	protected := api.Group("")
	protected.Use(middleware.Auth(tokens))
	protected.POST("/playground/run", middleware.RateLimit(cfg.PlaygroundRateMin), h.RunCode)
	protected.POST("/exercises/:id/submit", middleware.RateLimit(cfg.PlaygroundRateMin), h.SubmitExercise)
	protected.POST("/quizzes/:id/answer", h.AnswerQuiz)
	protected.POST("/lessons/:slug/complete", h.CompleteLesson)

	me := protected.Group("/me")
	me.GET("/progress", h.Progress)
	me.GET("/dashboard", h.Dashboard)
	me.GET("/review/today", h.ReviewToday)
	me.POST("/review/:flashcardId", h.ReviewCard)
	me.GET("/stats", h.Stats)
	me.GET("/notes", h.Notes)
	me.POST("/notes", h.SaveNote)
	me.PUT("/notes/:id", h.UpdateNote)
	me.DELETE("/notes/:id", h.DeleteNote)
	me.GET("/bookmarks", h.Bookmarks)
	me.POST("/bookmarks", h.AddBookmark)
	me.DELETE("/bookmarks/:lessonId", h.RemoveBookmark)
	protected.POST("/projects/steps/:id/complete", h.CompleteProjectStep)

	return r
}
