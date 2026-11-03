package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"golearn/internal/middleware"
	"golearn/internal/response"
	"golearn/internal/service"
)

type Handlers struct {
	Auth    *service.AuthService
	Content *service.ContentService
	Learn   *service.LearnService
	Me      *service.MeService
}

func fail(c *gin.Context, err error) {
	var appErr *service.Error
	if errors.As(err, &appErr) {
		response.Fail(c, appErr.Status, appErr.Code, appErr.Message)
		return
	}
	slog.Error("request failed", "path", c.FullPath(), "error", err)
	response.Fail(c, http.StatusInternalServerError, "internal", "Terjadi kesalahan pada server.")
}

func badBody(c *gin.Context, err error) {
	response.Fail(c, http.StatusUnprocessableEntity, "validation", "Body permintaan tidak valid: "+err.Error())
}

func (h *Handlers) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		badBody(c, err)
		return
	}
	res, err := h.Auth.Register(req.Name, req.Email, req.Password)
	if err != nil {
		fail(c, err)
		return
	}
	response.Created(c, res)
}

func (h *Handlers) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		badBody(c, err)
		return
	}
	res, err := h.Auth.Login(req.Email, req.Password)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, res)
}

func (h *Handlers) Refresh(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		badBody(c, err)
		return
	}
	res, err := h.Auth.Refresh(req.RefreshToken)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, res)
}

func (h *Handlers) Logout(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	_ = c.ShouldBindJSON(&req)
	if err := h.Auth.Logout(req.RefreshToken); err != nil {
		fail(c, err)
		return
	}
	response.OK(c, gin.H{"logged_out": true})
}

func (h *Handlers) CurrentUser(c *gin.Context) {
	user, stats, err := h.Auth.Me(middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, gin.H{"user": user, "stats": stats})
}

func (h *Handlers) DeleteAccount(c *gin.Context) {
	if err := h.Auth.DeleteAccount(middleware.UserID(c)); err != nil {
		fail(c, err)
		return
	}
	response.OK(c, gin.H{"deleted": true})
}
