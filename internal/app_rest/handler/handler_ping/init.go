package handler_ping

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helicopter-hr/config"
	"helicopter-hr/internal/app_rest/middleware/jwtx"
	"helicopter-hr/internal/app_rest/service/service_auth"
	"helicopter-hr/pkg/ginx"
	"net/http"
)

type Handler struct {
	cfg         *config.ConfigApp
	authService service_auth.AuthServiceInterface
	jwtAuth     jwtx.AuthenticationInterface
}

func NewHandlerPing(router *gin.Engine) {

	v1 := router.Group("/api/v1/")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			var (
				guid = ctx.Value("request_id").(string)
			)
			zap.L().Info("ping", zap.String("request_id", guid))
			ginx.RespondWithJSON(ctx, http.StatusOK, "pong", nil)
		})
	}
}
