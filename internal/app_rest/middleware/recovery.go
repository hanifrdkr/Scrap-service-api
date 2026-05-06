package middleware

import (
	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ErrorHandler(ctx *gin.Context, err any) {
	httpResponse := HttpResponse{Message: "Internal server error", Status: 500}
	ctx.AbortWithStatusJSON(500, httpResponse)
}
