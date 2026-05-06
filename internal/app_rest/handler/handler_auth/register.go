package handler_auth

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helicopter-hr/internal/app_rest/service/service_auth"
	"helicopter-hr/pkg/ginx"
	"helicopter-hr/pkg/validatorx"
	"net/http"
)

func (h *Handler) HandlerRegister(ctx *gin.Context) {
	var (
		guid  = ctx.Value("request_id").(string)
		param service_auth.RegisterPayload
	)

	cLogger := zap.L().With(
		zap.String("layer", "handler.register"),
		zap.String("request_id", guid),
	)

	if err := ctx.ShouldBindJSON(&param); err != nil {
		cLogger.Error("error decode payload register", zap.Error(err))
		ginx.RespondWithError(
			ctx,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil,
		)
		return
	}

	if err := validatorx.Validate(param); err != nil {
		cLogger.Warn("error validate payload register", zap.Error(err))
		ginx.RespondWithError(ctx, http.StatusUnprocessableEntity, err.Error(), validatorx.ExtractError(err))
	}

	err := h.authService.Register(ctx, param)
	if err != nil {
		switch err.Error() {
		case "this username already exists":
			cLogger.Error("error this username already exist", zap.Error(err))
			ginx.RespondWithJSON(ctx,
				http.StatusUnprocessableEntity,
				http.StatusText(http.StatusUnprocessableEntity),
				err.Error(),
			)
			return
		default:
			cLogger.Error("err", zap.Error(err))
			ginx.RespondWithJSON(ctx, http.StatusInternalServerError, "err", err.Error())
			return
		}

	}

	cLogger.Error("success register", zap.Error(err))
	ginx.RespondWithJSON(ctx, http.StatusOK, "success register", nil)
	return
}
