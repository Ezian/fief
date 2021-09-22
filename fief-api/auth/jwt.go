package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

const (
	// TODO put it in ENV
	mySecretKeyForJWT = "qsfq<fgzegqzegqzergrqegrqe"
)

func GenerateJWT(userLogin string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["login"] = userLogin
	claims["exp"] = time.Now().Add(time.Minute * 300).Unix()

	tokenString, err := token.SignedString([]byte(mySecretKeyForJWT))
	if err != nil {
		// TODO Wrap error in a correct way
		return "", err
	}
	return tokenString, nil
}

func ValidateHeader(bearerHeader string) (interface{}, error) {
	bearerToken := strings.Split(bearerHeader, " ")[1]
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// TODO Handle the error correctly
			return nil, fmt.Errorf("error decoding token")
		}
		return []byte(mySecretKeyForJWT), nil
	})
	if err != nil {
		// TODO Handle the error correctly
		return nil, err
	}
	if token.Valid {
		return claims["login"].(string), nil
	}
	return nil, errors.New("invalid token")
}
