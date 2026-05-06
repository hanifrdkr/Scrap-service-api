package repo_auth

import (
	"context"
	"helicopter-hr/internal/app_rest/model"
)

func (a *authRepository) FindOneAccessToken(ctx context.Context, p model.AccessToken) (*model.AccessToken, error) {
	return nil, nil
}
