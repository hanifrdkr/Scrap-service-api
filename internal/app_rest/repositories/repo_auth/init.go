package repo_auth

import (
	"context"
	"github.com/jmoiron/sqlx"
	"helicopter-hr/internal/app_rest/model"
)

type AuthRepositoryInterface interface {
	FindOne(ctx context.Context, p model.Auth) (result *model.Auth, err error)
	FindOneAccessToken(ctx context.Context, p model.AccessToken) (result *model.AccessToken, err error)
	StoreAuth(ctx context.Context, p model.Auth) (err error)
	StoreAccessToken(ctx context.Context, p model.AccessToken) (err error)
	UpdateOne(ctx context.Context, p model.AccessToken) (err error)
}

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) AuthRepositoryInterface {
	return &authRepository{
		db: db,
	}
}
