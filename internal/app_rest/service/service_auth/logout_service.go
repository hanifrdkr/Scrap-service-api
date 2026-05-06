package service_auth

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"helicopter-hr/internal/app_rest/model"
	"time"
)

func (a *authService) Logout(ctx context.Context, token string) error {
	var (
		guid   = ctx.Value("request_id").(string)
		userID = ctx.Value("uid").(string)
	)

	cLogger := zap.L().With(
		zap.String("layer", "service.logout"),
		zap.String("request_id", guid),
	)

	err := a.repoAuth.UpdateOne(ctx, model.AccessToken{
		IsRevoke:  true,
		UpdatedAt: time.Now(),
		UserID:    userID,
	})
	if err != nil {
		cLogger.Error("error update by token", zap.Error(err))
		err = errors.New("username or password is incorrect")
		return err
	}

	cLogger.Info("success logout service")
	return nil
}
