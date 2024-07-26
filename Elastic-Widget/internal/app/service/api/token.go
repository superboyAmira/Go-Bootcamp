package api

import (
	"errors"
	"goday03/src/internal/app/service/response"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var key []byte = []byte("key")

type TokenJWT struct {
	Token string
}

func (token *TokenJWT) Generate(w http.ResponseWriter, req *http.Request) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	JWTToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var err error
	token.Token, err = JWTToken.SignedString(key)
	if err != nil {
		resp := &response.TokenHTTP{
			Err: err.Error(),
		}
		resp.SendError(w)
		return
	}

	resp := &response.TokenHTTP{
		Token: token.Token,
	}
	resp.SendResponse(w)
}

func (token *TokenJWT) Validate(w http.ResponseWriter, req *http.Request) bool {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		resp := &response.TokenHTTP{
			Err: "missing Authorization header",
		}
		resp.SendUnauthorized(w)
		return false
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		resp := &response.TokenHTTP{
			Err: "invalid Authorization header format",
		}
		resp.SendUnauthorized(w)
		return false
	}

	token.Token = parts[1]

	Parsedtoken, err := jwt.Parse(token.Token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return key, nil
	})

	if err != nil {
		resp := &response.TokenHTTP{
			Err: err.Error(),
		}
		resp.SendError(w)
		return false
	}

	if !Parsedtoken.Valid {
		resp := &response.TokenHTTP{
			Err: "invalid token",
		}
		resp.SendUnauthorized(w)
		return false
	}

	// resp := &response.TokenHTTP{
	// 	Token: token.Token,
	// }
	// resp.SendResponse(w)
	return true
}
