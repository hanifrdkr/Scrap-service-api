package repo_user

import (
	"context"
	"github.com/jmoiron/sqlx"
	"helicopter-hr/internal/app_rest/model"
)

type UserRepositoryInterface interface {
	FindOne(ctx context.Context, p model.User) (*model.User, error)
	StoreUser(ctx context.Context, p model.User) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}
