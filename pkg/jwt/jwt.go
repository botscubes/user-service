package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(id int, key string) (string, error) {
	claims := UserClaims{
		id,
		jwt.RegisteredClaims{},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(key))
	return s, err
}

func GetIdFromToken(tokenString string, key string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims.Id, nil
	}

	return 0, errors.New("token is not valid")
}
