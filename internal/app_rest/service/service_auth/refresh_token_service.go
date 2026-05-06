package service_auth

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"go.uber.org/zap"
	"helicopter-hr/internal/app_rest/model"
	"time"
)

func (a *authService) RefreshToken(ctx context.Context, rToken string) (*RefreshTokenResponse, error) {
	var (
		guid   = ctx.Value("request_id").(string)
		result RefreshTokenResponse
	)

	cLogger := zap.L().With(
		zap.String("layer", "service.logout"),
		zap.String("request_id", guid),
	)

	token, err := a.jwtAuth.ParseRefreshToken(rToken)
	if err != nil {
		cLogger.Error("err parse jwt refresh token", zap.Error(err))
		err = errors.New("unauthorized")
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if !ok && !token.Valid {
		cLogger.Error("err claim refresh token", zap.Error(err))
		err = errors.New("unauthorized")
		return nil, err
	}

	userID, ok := claims["UserID"].(string) //convert the interface to string
	if !ok {
		cLogger.Error("err get claims user id", zap.Error(err))
		err = errors.New("unauthorized")
		return nil, err
	}

	auth, err := a.repoAuth.FindOne(ctx, model.Auth{UserID: userID})
	if err != nil {
		cLogger.Error("error find auth by user id", zap.Error(err))
		return nil, err
	}

	if auth == nil {
		err = errors.New("user not found")
		return nil, err
	}

	newToken, newRefreshToken, _ := a.jwtAuth.GenerateAllTokens(
		userID,
	)

	accessToken := model.AccessToken{
		Token:        newToken,
		RefreshToken: newRefreshToken,
		IsRevoke:     false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		ExpiredAt:    time.Now().Local().Add(time.Hour * time.Duration(168)),
		UserID:       auth.UserID,
	}

	err = a.repoAuth.StoreAccessToken(ctx, accessToken)
	if err != nil {
		cLogger.Error("error store access token", zap.Error(err))
		return nil, err
	}

	result.Token = newToken
	result.RefreshToken = newRefreshToken

	cLogger.Info("success service refresh token", zap.Error(err))
	return &result, nil
}
