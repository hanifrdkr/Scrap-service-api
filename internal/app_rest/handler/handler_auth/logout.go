package handler_auth

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helicopter-hr/pkg/ginx"
	"net/http"
)

func (h *Handler) HandlerLogout(ctx *gin.Context) {
	var (
		guid  = ctx.Value("request_id").(string)
		token = ctx.GetString("token")
	)

	cLogger := zap.L().With(
		zap.String("layer", "handler.logout"),
		zap.String("request_id", guid),
	)

	err := h.authService.Logout(ctx, token)
	if err != nil {
		cLogger.Error("error logout", zap.Error(err))
		ginx.RespondWithJSON(ctx, http.StatusInternalServerError, "err", err.Error())
		return
	}

	cLogger.Info("success logout", zap.Error(err))
	ginx.RespondWithJSON(ctx, http.StatusOK, "success", nil)
	return
}
