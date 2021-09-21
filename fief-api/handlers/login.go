package handlers

import (
	"fief/auth"
	"fief/models"
	"fief/restapi/operations/user"

	"github.com/boltdb/bolt"
	"github.com/go-openapi/runtime/middleware"
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
	ok, err := auth.CheckUserPassword(impl.dbClient, *params.Login.Login, *params.Login.Password)
	if err != nil {
		return user.NewLoginInternalServerError().WithPayload("Error fetching user details")
	}

	if !ok {
		return user.NewRegisterNotFound()
	}

	token, err := auth.GenerateJWT(*params.Login.Login)
	if err != nil {
		return user.NewLoginInternalServerError().WithPayload("Error defining token")
	}
	return user.NewLoginOK().WithPayload(&models.LoginSuccess{Success: true, Token: token})
}
