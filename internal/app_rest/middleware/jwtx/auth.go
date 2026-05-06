package jwtx

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helicopter-hr/internal/app_rest/model"
	"helicopter-hr/pkg/ginx"
	"net/http"
	"strings"
)

func (j *AuthenticationJWT) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			guid = ctx.Value("request_id").(string)
		)

		cLogger := zap.L().With(
			zap.String("layer", "service.login"),
			zap.String("request_id", guid),
		)

		authorization := ctx.Request.Header.Get("Authorization")
		clientToken := strings.TrimPrefix(authorization, "Bearer ")
		if clientToken == "" {
			ginx.RespondWithJSON(
				ctx,
				http.StatusUnauthorized,
				"No Authorization header provided",
				nil,
			)
			return
		}

		existingToken, err := j.repoAuth.FindOneAccessToken(ctx, model.AccessToken{Token: clientToken, IsRevoke: true})
		if err != nil {
			cLogger.Error("err find one access token", zap.Error(err))
			ginx.RespondWithJSON(
				ctx,
				http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
				nil,
			)
			return
		}

		if existingToken != nil && existingToken.IsRevoke {
			ginx.RespondWithJSON(
				ctx,
				http.StatusUnauthorized,
				"Something wrong with authorization",
				nil,
			)
			return
		}

		claims, msg := j.ValidateToken(clientToken)
		if msg != "" {
			ginx.RespondWithJSON(
				ctx,
				http.StatusUnauthorized,
				"Invalid Authorization",
				nil,
			)
			return
		}

		ctx.Set("username", claims.Username)
		ctx.Set("token", clientToken)
		ctx.Set("uid", claims.UserID)

		ctx.Next()
	}
}
