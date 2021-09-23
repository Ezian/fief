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
		return "", errors.Wrap(err, "cannot create JWT signed token")
	}
	return tokenString, nil
}

func ValidateHeader(bearerHeader string) (interface{}, error) {
	bearerToken := strings.Split(bearerHeader, " ")[1]
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("mismatching signing method")
		}
		return []byte(mySecretKeyForJWT), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "cannot decode the token")
	}
	if token.Valid {
		return claims["login"].(string), nil
	}
	return nil, errors.New("invalid token")
}
