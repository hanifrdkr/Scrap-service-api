package jwtx

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

type SignedDetails struct {
	Username string
	UserID   string
	jwt.StandardClaims
}

func (j *AuthenticationJWT) GenerateAllTokens(userID string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(time.Now().UTC().Add(24 * time.Hour).Unix())),
		},
	}

	refreshClaims := &SignedDetails{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(time.Now().UTC().Add(48 * time.Hour).Unix())),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.cfg.App.SecretKey))
	if err != nil {
		return
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(j.cfg.App.SecretKey))
	if err != nil {
		return
	}

	return token, refreshToken, err
}

func (j *AuthenticationJWT) ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.cfg.App.SecretKey), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}

	return claims, msg
}

func (j *AuthenticationJWT) ParseRefreshToken(rToken string) (*jwt.Token, error) {
	t, err := jwt.Parse(rToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.cfg.App.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return t, nil
}
