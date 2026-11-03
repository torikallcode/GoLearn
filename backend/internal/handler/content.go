package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"golearn/internal/middleware"
	"golearn/internal/response"
)

func (h *Handlers) Paths(c *gin.Context) {
	paths, err := h.Content.Paths()
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, paths)
}

func (h *Handlers) Path(c *gin.Context) {
	detail, err := h.Content.PathDetail(c.Param("slug"), middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, detail)
}

func (h *Handlers) Lesson(c *gin.Context) {
	detail, err := h.Content.Lesson(c.Param("slug"), middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, detail)
}

func (h *Handlers) LessonExercises(c *gin.Context) {
	items, err := h.Content.Exercises(c.Param("slug"), middleware.UserID(c))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, items)
}

func (h *Handlers) LessonQuizzes(c *gin.Context) {
	items, err := h.Content.Quizzes(c.Param("slug"))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, items)
}

func (h *Handlers) Search(c *gin.Context) {
	hits, err := h.Content.Search(c.Query("q"))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, hits)
}

func (h *Handlers) RunCode(c *gin.Context) {
	var req struct {
		Code string `json:"code" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		badBody(c, err)
		return
	}
	res, err := h.Learn.Run(c.Request.Context(), req.Code)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, res)
}

func (h *Handlers) SubmitExercise(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, 400, "bad_id", "ID latihan tidak valid.")
		return
	}
	var req struct {
		Code string `json:"code" binding:"required"`
	}
	if err = c.ShouldBindJSON(&req); err != nil {
		badBody(c, err)
		return
	}
	res, err := h.Learn.SubmitExercise(c.Request.Context(), middleware.UserID(c), id, req.Code)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, res)
}

func (h *Handlers) AnswerQuiz(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, 400, "bad_id", "ID kuis tidak valid.")
		return
	}
	var req struct {
		OptionID *uint64 `json:"option_id"`
		Answer   string  `json:"answer"`
	}
	if err = c.ShouldBindJSON(&req); err != nil {
		badBody(c, err)
		return
	}
	res, err := h.Learn.AnswerQuiz(middleware.UserID(c), id, req.OptionID, req.Answer)
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, res)
}

func (h *Handlers) CompleteLesson(c *gin.Context) {
	res, err := h.Learn.CompleteLesson(middleware.UserID(c), c.Param("slug"))
	if err != nil {
		fail(c, err)
		return
	}
	response.OK(c, res)
}
