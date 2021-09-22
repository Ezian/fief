package handlers

import (
	"fief/auth"
	"fief/models"
	"fief/restapi/operations/user"

	"github.com/boltdb/bolt"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
)

type loginImpl struct {
	dbClient *bolt.DB
}

func NewUserLoginHandler(db *bolt.DB) user.LoginHandler {
	return &loginImpl{
		dbClient: db,
	}
}

func (impl *loginImpl) Handle(params user.LoginParams) middleware.Responder {
	log := log.WithField("login", *params.Login.Login)
	ok, err := auth.CheckUserPassword(impl.dbClient, *params.Login.Login, *params.Login.Password)
	if err != nil {
		log.WithError(err).Error("Error fetching user details")
		return user.NewLoginInternalServerError().WithPayload("Error fetching user details")
	}

	if !ok {
		log.Warn("Wrong login/password")
		return user.NewLoginUnauthorized().WithPayload("Wrong login/password")
	}

	token, err := auth.GenerateJWT(*params.Login.Login)
	if err != nil {
		log.WithError(err).Error("Error when generating the token")
		return user.NewLoginInternalServerError().WithPayload("Error when generating the token")
	}
	log.Info("Login Successful")
	return user.NewLoginOK().WithPayload(&models.LoginSuccess{Success: true, Token: token})
}
