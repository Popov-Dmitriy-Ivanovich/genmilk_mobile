package user

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var ACCESS_EXP = time.Now().Add(time.Hour * 24)
var REFRESH_EXP = time.Now().AddDate(50, 0, 1)

type BaseToken struct {
	jwt.RegisteredClaims
	UserID uint
}

type AccessToken struct {
	BaseToken
}

type RefreshToken struct {
	BaseToken
}

func GenerateTokens(UserID uint) (string, string, error) {
	accessTokenClaims := AccessToken{
		BaseToken: BaseToken{
			UserID: UserID,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(ACCESS_EXP),
			},
		},
	}
	refreshTokenClaims := RefreshToken{
		BaseToken: BaseToken{
			UserID: UserID,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(REFRESH_EXP),
			},
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", "", err
	}
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", "", err
	}
	return accessTokenString, refreshTokenString, nil
}
