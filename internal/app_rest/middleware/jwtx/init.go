package jwtx

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"helicopter-hr/config"
	"helicopter-hr/internal/app_rest/repositories/repo_auth"
)

type AuthenticationJWT struct {
	cfg      *config.ConfigApp
	repoAuth repo_auth.AuthRepositoryInterface
}
type AuthenticationInterface interface {
	Authentication() gin.HandlerFunc
	GenerateAllTokens(userID string) (signedToken string, signedRefreshToken string, err error)
	ValidateToken(signedToken string) (claims *SignedDetails, msg string)
	ParseRefreshToken(rToken string) (*jwt.Token, error)
}

func NewJWTAuthentication(
	cfg *config.ConfigApp,
	repoAuth repo_auth.AuthRepositoryInterface) AuthenticationInterface {
	return &AuthenticationJWT{
		cfg:      cfg,
		repoAuth: repoAuth,
	}
}
