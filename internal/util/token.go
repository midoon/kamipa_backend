package util

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/midoon/kamipa_backend/internal/domain"
	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
	"github.com/midoon/kamipa_backend/internal/helper"
)

type TokenUtil struct {
	SecretKey string
	RedisRepo domain.RedisRepository
}

func NewTokenUtil(secretKey string, redisRepo domain.RedisRepository) *TokenUtil {
	return &TokenUtil{
		SecretKey: secretKey,
		RedisRepo: redisRepo,
	}
}

func (t TokenUtil) CreateToken(ctx context.Context, user *kamipa_entity.User, exp int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user.ID,
		"expire": exp,
	})

	jwtToken, err := token.SignedString([]byte(t.SecretKey))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (t TokenUtil) ParseToken(ctx context.Context, jwtToken string) (string, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.SecretKey), nil
	})

	if err != nil {
		return "", helper.NewCustomError(http.StatusUnauthorized, "Unauthorized", err)
	}

	claims := token.Claims.(jwt.MapClaims)

	expire := claims["expire"].(float64)
	if int64(expire) < time.Now().UnixMilli() {
		return "", helper.NewCustomError(http.StatusUnauthorized, "Token expired", nil)
	}

	userId := claims["id"].(string)
	redisKey := fmt.Sprintf("refresh_token:%s", userId)

	result, err := t.RedisRepo.ExistData(ctx, redisKey)
	if err != nil {
		return "", helper.NewCustomError(http.StatusInternalServerError, "redis error", err)
	}

	if result == 0 {
		return "", helper.NewCustomError(http.StatusUnauthorized, "Unauthorized", err)
	}

	return userId, nil
	// token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	// 	}
	// 	return []byte(t.SecretKey), nil
	// })

	// if err != nil {
	// 	return "", err
	// }
	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	if id, ok := claims["id"].(string); ok {
	// 		return id, nil
	// 	}

	// 	return "", fmt.Errorf("id claim not found or invalid")
	// }

	// return "", fmt.Errorf("invalid token")
}
