package service_auth

import (
	"context"
	"helicopter-hr/internal/app_rest/middleware/jwtx"
	"helicopter-hr/internal/app_rest/repositories/repo_auth"
	"helicopter-hr/internal/app_rest/repositories/repo_user"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, l LoginPayload) (result *LoginResponse, err error)
	Register(ctx context.Context, l RegisterPayload) (err error)
	Profile(ctx context.Context, userID string, username string) (*ProfileReturn, error)
	Logout(ctx context.Context, token string) error
	RefreshToken(ctx context.Context, refreshToken string) (*RefreshTokenResponse, error)
}

type authService struct {
	repoAuth repo_auth.AuthRepositoryInterface
	repoUser repo_user.UserRepositoryInterface
	jwtAuth  jwtx.AuthenticationInterface
}

func NewAuthService(repoAuth repo_auth.AuthRepositoryInterface, repoUser repo_user.UserRepositoryInterface, jwtAuth jwtx.AuthenticationInterface) AuthServiceInterface {
	return &authService{
		repoAuth: repoAuth,
		repoUser: repoUser,
		jwtAuth:  jwtAuth,
	}
}
