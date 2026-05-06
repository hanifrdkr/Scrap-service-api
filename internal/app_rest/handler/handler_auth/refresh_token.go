package handler_auth

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helicopter-hr/pkg/ginx"
	"net/http"
)

func (h *Handler) HandlerRefreshToken(ctx *gin.Context) {
	var (
		guid         = ctx.Value("request_id").(string)
		refreshToken = ctx.Request.Header.Get("Refresh-Token")
	)

	cLogger := zap.L().With(
		zap.String("layer", "handler.refresh_token"),
		zap.String("request_id", guid),
	)

	if refreshToken == "" {
		ginx.RespondWithJSON(ctx, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return
	}

	result, err := h.authService.RefreshToken(ctx, refreshToken)
	if err != nil {
		switch err.Error() {
		case "unauthorized":
			cLogger.Error("error unauthorized", zap.Error(err))
			ginx.RespondWithError(
				ctx,
				http.StatusUnauthorized,
				http.StatusText(http.StatusUnauthorized),
				err.Error(),
			)
			return
		default:
			cLogger.Error("err ", zap.Error(err))
			ginx.RespondWithError(
				ctx,
				http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
				err.Error(),
			)
			return
		}
	}

	cLogger.Info("success handler refresh token")
	ginx.RespondWithJSON(ctx, http.StatusOK, "success", result)
}
