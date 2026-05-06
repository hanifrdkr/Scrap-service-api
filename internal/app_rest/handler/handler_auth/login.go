package handler_auth

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helicopter-hr/internal/app_rest/service/service_auth"
	"helicopter-hr/pkg/ginx"
	"helicopter-hr/pkg/validatorx"
	"net/http"
)

func (h *Handler) HandlerLogin(ctx *gin.Context) {
	var (
		guid  = ctx.Value("request_id").(string)
		param service_auth.LoginPayload
	)

	cLogger := zap.L().With(
		zap.String("layer", "handler.login"),
		zap.String("request_id", guid),
	)

	if err := ctx.ShouldBindJSON(&param); err != nil {
		cLogger.Error("error decode payload login", zap.Error(err))
		ginx.RespondWithError(
			ctx,
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err,
		)
		return
	}

	if err := validatorx.Validate(param); err != nil {
		cLogger.Warn("error validate payload login", zap.Error(err))
		ginx.RespondWithError(ctx, http.StatusUnprocessableEntity, err.Error(), validatorx.ExtractError(err))
		return
	}

	result, err := h.authService.Login(ctx, param)
	if err != nil {
		switch err.Error() {
		case "username or password is incorrect":
			cLogger.Warn("username or password is incorrect", zap.Error(err))
			ginx.RespondWithJSON(
				ctx,
				http.StatusUnauthorized,
				http.StatusText(http.StatusUnauthorized),
				err.Error(),
			)
			return
		case "unauthorized":
			cLogger.Warn("unauthorized", zap.Error(err))
			ginx.RespondWithJSON(
				ctx,
				http.StatusUnauthorized,
				http.StatusText(http.StatusUnauthorized),
				err.Error(),
			)
			return
		default:
			cLogger.Warn("err ", zap.Error(err))
			ginx.RespondWithJSON(
				ctx,
				http.StatusInternalServerError,
				http.StatusText(http.StatusUnauthorized),
				err.Error(),
			)
			return
		}
	}

	cLogger.Info("success handler login")
	ginx.RespondWithJSON(ctx, http.StatusOK, "success", result)
	return
}
