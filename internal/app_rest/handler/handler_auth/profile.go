package handler_auth

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helicopter-hr/pkg/ginx"
	"net/http"
)

func (h *Handler) HandlerProfile(ctx *gin.Context) {
	var (
		guid     = ctx.Value("request_id").(string)
		userID   = ctx.GetString("user_id")
		username = ctx.GetString("username")
	)

	cLogger := zap.L().With(
		zap.String("layer", "handler.logout"),
		zap.String("request_id", guid),
	)

	result, err := h.authService.Profile(ctx, userID, username)
	if err != nil {
		cLogger.Error("error get profile", zap.Error(err))
		ginx.RespondWithJSON(ctx, http.StatusInternalServerError, "err", err.Error())
		return
	}

	cLogger.Info("success get profile")
	ginx.RespondWithJSON(ctx, http.StatusOK, "success", result)
	return
}
