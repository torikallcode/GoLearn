package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data, "error": nil, "meta": nil})
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, gin.H{"success": true, "data": data, "error": nil, "meta": nil})
}

func Meta(c *gin.Context, data any, meta any) {
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data, "error": nil, "meta": meta})
}

func Fail(c *gin.Context, status int, code, message string) {
	c.AbortWithStatusJSON(status, gin.H{
		"success": false,
		"data":    nil,
		"error":   ErrorBody{Code: code, Message: message},
		"meta":    nil,
	})
}
