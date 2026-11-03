package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"golearn/internal/middleware"
	"golearn/internal/response"
)

func (h *Handlers) Dashboard(c *gin.Context) {
	out, err := h.Me.Dashboard(middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, out)
}

func (h *Handlers) Progress(c *gin.Context) {
	out, err := h.Me.Progress(middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, out)
}

func (h *Handlers) Stats(c *gin.Context) {
	out, err := h.Me.Stats(middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, out)
}

func (h *Handlers) ReviewToday(c *gin.Context) {
	out, err := h.Me.ReviewToday(middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, out)
}

func (h *Handlers) ReviewCard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("flashcardId"), 10, 64)
	if err != nil {
		response.Fail(c, 400, "bad_id", "ID flashcard tidak valid.")
		return
	}
	var req struct {
		Quality int `json:"quality"`
	}
	if err = c.ShouldBindJSON(&req); err != nil {
		badBody(c, err)
		return
	}
	out, err := h.Me.ReviewCard(middleware.UserID(c), id, req.Quality)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, out)
}

func (h *Handlers) Notes(c *gin.Context) {
	var lessonID uint64
	if raw := c.Query("lesson_id"); raw != "" {
		lessonID, _ = strconv.ParseUint(raw, 10, 64)
	}
	out, err := h.Me.Notes(middleware.UserID(c), lessonID)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, out)
}

func (h *Handlers) SaveNote(c *gin.Context) {
	var req struct {
		LessonID uint64 `json:"lesson_id" binding:"required"`
		Content  string `json:"content"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		badBody(c, err)
		return
	}
	out, err := h.Me.UpsertNote(middleware.UserID(c), req.LessonID, req.Content)
	if err != nil {
		fail(c, err)
		return
	}
	response.Created(c, out)
}

func (h *Handlers) UpdateNote(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, 400, "bad_id", "ID catatan tidak valid.")
		return
	}
	var req struct {
		Content string `json:"content"`
	}
	if err = c.ShouldBindJSON(&req); err != nil {
		badBody(c, err)
		return
	}
	if err = h.Me.UpdateNote(middleware.UserID(c), id, req.Content); err != nil {
		fail(c, err)
		return
	}
	response.OK(c, gin.H{"updated": true})
}

func (h *Handlers) DeleteNote(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, 400, "bad_id", "ID catatan tidak valid.")
		return
	}
	if err = h.Me.DeleteNote(middleware.UserID(c), id); err != nil {
		fail(c, err)
		return
	}
	response.OK(c, gin.H{"deleted": true})
}

func (h *Handlers) Bookmarks(c *gin.Context) {
	out, err := h.Me.Bookmarks(middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, out)
}

func (h *Handlers) AddBookmark(c *gin.Context) {
	var req struct {
		LessonID uint64 `json:"lesson_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		badBody(c, err)
		return
	}
	if err := h.Me.AddBookmark(middleware.UserID(c), req.LessonID); err != nil {
		fail(c, err)
		return
	}
	response.Created(c, gin.H{"bookmarked": true})
}

func (h *Handlers) RemoveBookmark(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("lessonId"), 10, 64)
	if err != nil {
		response.Fail(c, 400, "bad_id", "ID pelajaran tidak valid.")
		return
	}
	if err = h.Me.RemoveBookmark(middleware.UserID(c), id); err != nil {
		fail(c, err)
		return
	}
	response.OK(c, gin.H{"bookmarked": false})
}

func (h *Handlers) Projects(c *gin.Context) {
	out, err := h.Me.Projects(middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, out)
}

func (h *Handlers) Project(c *gin.Context) {
	out, err := h.Me.Project(c.Param("slug"), middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, out)
}

func (h *Handlers) CompleteProjectStep(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, 400, "bad_id", "ID langkah tidak valid.")
		return
	}
	if err = h.Me.CompleteProjectStep(middleware.UserID(c), id); err != nil {
		fail(c, err)
		return
	}
	response.OK(c, gin.H{"done": true})
}
