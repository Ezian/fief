package handlers

import (
	"fief/auth"
	"fief/models"
	"fief/restapi/operations/user"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/go-openapi/runtime/middleware"
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
	err := auth.RegisterNewUser(impl.dbClient, params.Signup)
	if err != nil {
		// TODO log error before return
		return user.NewRegisterInternalServerError().WithPayload(fmt.Sprintf("Error registering user: %s", err))
	}
	return user.NewRegisterOK().WithPayload(&models.SuccessResponse{Success: true, Message: "User Registered successfully"})

}
