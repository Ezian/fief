package handlers

import (
	"errors"
	"fief/auth"
	"fief/restapi/operations/user"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
)

type registerImpl struct {
	dbClient *bolt.DB
}

func NewUserRegisterHandler(db *bolt.DB) user.RegisterHandler {
	return &registerImpl{
		dbClient: db,
	}
}

func (impl *registerImpl) Handle(params user.RegisterParams) middleware.Responder {
	log := log.WithField("login", *params.Signup.Login).WithField("email", *params.Signup.Email)
	err := auth.RegisterNewUser(impl.dbClient, params.Signup)
	if err != nil {
		log.WithError(err).Error("Error registering user")
		if errors.Is(err, auth.UserAlreadyExists) {
			return user.NewRegisterBadRequest().WithPayload(fmt.Sprintf("Error registering user: %s (User already exists)", *params.Signup.Login))
		}
		return user.NewRegisterInternalServerError().WithPayload(fmt.Sprintf("Cannot register user: %s", *params.Signup.Login))
	}
	log.Info("Register user successful")
	return user.NewRegisterOK()
}
