package ginx

import (
	"github.com/gin-gonic/gin"
)

type generalResponse struct {
	Message string      `json:"message,omitempty"`
	Status  bool        `json:"status,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorField struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func RespondWithError(ctx *gin.Context, status int, message string, error interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, generalResponse{
		Status:  getStatusBool(status),
		Message: message,
		Error:   error,
	})
	ctx.Abort()
	return
}

func RespondWithJSON(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, generalResponse{
		Status:  getStatusBool(status),
		Message: message,
		Data:    data,
	})
	ctx.Abort()
	return
}

func getStatusBool(status int) bool {
	switch status {
	case 400, 404, 422, 401, 500:
		return false
	default:
		return true
	}
}
