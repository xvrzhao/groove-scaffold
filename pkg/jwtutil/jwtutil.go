package jwtutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secret string

func init() {
	secret = os.Getenv("JWT_SECRET")
	fmt.Println("secret:", secret)
}

type TokenPayload struct {
	UserID   int    `json:"userId"`
	Username string `json:"username"`
}

func (tp *TokenPayload) retrieveFromClaims(mapClaims jwt.MapClaims) {
	byts, _ := json.Marshal(mapClaims)
	json.Unmarshal(byts, tp)
}

type tokenClaims struct {
	TokenPayload
	jwt.RegisteredClaims
}

func Gen(payload TokenPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		TokenPayload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Date(2099, 12, 31, 23, 59, 59, 0, time.Local)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	str, err := token.SignedString([]byte(secret))
	if err != nil {
		err = fmt.Errorf("token.SignedString failed: %w", err)
	}
	return str, err
}

func Parse(token string) (payload *TokenPayload, msg string, ok bool) {
	if token == "" {
		return nil, "token 不存在", false
	}

	tk, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if tk.Valid {
		if mapClaims, ok := tk.Claims.(jwt.MapClaims); ok {
			payload = new(TokenPayload)
			payload.retrieveFromClaims(mapClaims)
			return payload, "", ok
		} else {
			return nil, "解析 claims 失败", false
		}
	}

	switch {
	case errors.Is(err, jwt.ErrTokenMalformed) || errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return nil, "token 不合法", false
	case errors.Is(err, jwt.ErrTokenExpired):
		return nil, "token 已过期", false
	default:
		return nil, fmt.Sprintf("token 错误：%v", err), false
	}
}

func GetPayload(c *gin.Context) *TokenPayload {
	v, _ := c.Get("token")
	return v.(*TokenPayload)
}
