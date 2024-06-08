package token

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vakhia/artilight/internal/common/config"
	"time"
)

type IJwtService interface {
	GenerateToken(email string, userId string) (string, error)
	VerifyToken(token string) (string, error)
}

type JwtService struct {
	cfg *config.Config
}

func NewJWTService(cfg *config.Config) *JwtService {
	return &JwtService{
		cfg: cfg,
	}
}

func (j *JwtService) GenerateToken(email string, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(j.cfg.Token.Secret))
}

func (j *JwtService) VerifyToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(j.cfg.Token.Secret), nil
	})

	if err != nil {
		fmt.Println("Could not parse token")
		return "", errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return "", errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New("invalid token claims")
	}

	userId := claims["userId"].(string)
	return userId, nil
}
