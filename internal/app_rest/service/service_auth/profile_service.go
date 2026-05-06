package service_auth

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"helicopter-hr/internal/app_rest/model"
)

func (a *authService) Profile(ctx context.Context, userID string, username string) (*ProfileReturn, error) {
	var (
		guid   = ctx.Value("request_id").(string)
		result ProfileReturn
	)

	cLogger := zap.L().With(
		zap.String("layer", "service.profile"),
		zap.String("request_id", guid),
	)

	auth, err := a.repoAuth.FindOne(ctx, model.Auth{Email: username, UserID: userID})
	if err != nil {
		cLogger.Error("error find auth by username", zap.Error(err))
		err = errors.New("username or password is incorrect")
		return nil, err
	}

	result.Username = auth.Email

	cLogger.Info("success service profile")
	return &result, nil
}
