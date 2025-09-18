package util

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
	"github.com/redis/go-redis/v9"
)

type TokenUtil struct {
	SecretKey string
	Redis     *redis.Client
}

func NewTokenUtil(secretKey string, redisClient *redis.Client) *TokenUtil {
	return &TokenUtil{
		SecretKey: secretKey,
		Redis:     redisClient,
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

func (t TokenUtil) ParseToken(ctx context.Context, jwtToken string) (*kamipa_entity.User, error) {
	// token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(t.SecretKey), nil
	// })
	// if err != nil {
	// 	return nil, fiber.ErrUnauthorized
	// }

	// claims := token.Claims.(jwt.MapClaims)

	// expire := claims["expire"].(float64)
	// if int64(expire) < time.Now().UnixMilli() {
	// 	return nil, fiber.ErrUnauthorized
	// }

	// result, err := t.Redis.Exists(ctx, jwtToken).Result()
	// if err != nil {
	// 	return nil, err
	// }

	// if result == 0 {
	// 	return nil, fiber.ErrUnauthorized
	// }

	// id := claims["id"].(string)
	// auth := &model.Auth{
	// 	ID: id,
	// }
	return &kamipa_entity.User{}, nil
}
