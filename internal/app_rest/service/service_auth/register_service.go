package service_auth

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"helicopter-hr/internal/app_rest/model"
	"time"
)

func (a *authService) Register(ctx context.Context, l RegisterPayload) error {
	var (
		guid = ctx.Value("request_id").(string)
	)

	cLogger := zap.L().With(
		zap.String("layer", "service.signup"),
		zap.String("request_id", guid),
	)

	existingAuth, err := a.repoAuth.FindOne(ctx, model.Auth{Email: l.Email})
	if err != nil && err.Error() != "sql: no rows in result set" {
		cLogger.Error("err find one auth", zap.Error(err))
		return err
	}

	if existingAuth != nil {
		return errors.New("this email already exists")
	}

	password, err := HashPassword(l.Password)
	if err != nil {
		cLogger.Error("err hash password", zap.Error(err))
		return errors.New("err your password")
	}

	user := model.User{
		ID:        uuid.New().String(),
		Name:      l.Name,
		Email:     l.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = a.repoUser.StoreUser(ctx, user)
	if err != nil {
		cLogger.Error("err store auth", zap.Error(err))
		return err
	}

	auth := model.Auth{
		UserID:    user.ID,
		Password:  password,
		Email:     l.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	auth.Password = password

	err = a.repoAuth.StoreAuth(ctx, auth)
	if err != nil {
		cLogger.Error("err store auth", zap.Error(err))
		return err
	}

	cLogger.Error("success service register")
	return err
}

// HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", nil
	}

	return string(bytes), nil
}
