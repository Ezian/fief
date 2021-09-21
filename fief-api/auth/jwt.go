package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
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
