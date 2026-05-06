package service_auth

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"helicopter-hr/internal/app_rest/model"
	"time"
)

func (a *authService) Login(ctx context.Context, l LoginPayload) (*LoginResponse, error) {
	var (
		guid   = ctx.Value("request_id").(string)
		result LoginResponse
	)
	cLogger := zap.L().With(
		zap.String("layer", "service.login"),
		zap.String("request_id", guid),
	)

	auth, err := a.repoAuth.FindOne(ctx, model.Auth{Email: l.Email})
	if err != nil {
		cLogger.Error("error find auth by email", zap.Error(err))
		err = errors.New("username or password is incorrect")
		return nil, err
	}

	passwordIsValid, err := VerifyPassword(l.Password, auth.Password)
	if passwordIsValid != true || err != nil {
		cLogger.Error("error verify password", zap.Error(err))
		err = errors.New("username or password is incorrect")
		return nil, err
	}

	if auth == nil {
		return nil, err
	}
	token, refreshToken, _ := a.jwtAuth.GenerateAllTokens(
		auth.UserID,
	)

	accessToken := model.AccessToken{
		Token:        token,
		RefreshToken: refreshToken,
		IsRevoke:     false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		ExpiredAt:    time.Now().Local().Add(time.Hour * 24),
		UserID:       auth.UserID,
	}

	err = a.repoAuth.StoreAccessToken(ctx, accessToken)
	if err != nil {
		cLogger.Error("error", zap.Error(err))
		return nil, err
	}

	result.Token = token

	cLogger.Info("success service login")
	return &result, nil
}

func VerifyPassword(userPassword string, hashedPassword string) (bool, error) {
	check := true
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))
	if err != nil {
		check = false
		return false, err
	}

	return check, err
}
