// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"fief/restapi/operations"
	"fief/restapi/operations/game"
	"fief/restapi/operations/manage"
	"fief/restapi/operations/user"
)

//go:generate swagger generate server --target ../../fief-api --name Fief --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.FiefAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.FiefAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.GameGetGamesHandler == nil {
		api.GameGetGamesHandler = game.GetGamesHandlerFunc(func(params game.GetGamesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation game.GetGames has not yet been implemented")
		})
	}
	if api.GameGetGamesIDInstructionsHandler == nil {
		api.GameGetGamesIDInstructionsHandler = game.GetGamesIDInstructionsHandlerFunc(func(params game.GetGamesIDInstructionsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation game.GetGamesIDInstructions has not yet been implemented")
		})
	}
	if api.GameGetGamesIDStatusHandler == nil {
		api.GameGetGamesIDStatusHandler = game.GetGamesIDStatusHandlerFunc(func(params game.GetGamesIDStatusParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation game.GetGamesIDStatus has not yet been implemented")
		})
	}
	if api.UserLoginHandler == nil {
		api.UserLoginHandler = user.LoginHandlerFunc(func(params user.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation user.Login has not yet been implemented")
		})
	}
	if api.GamePostGamesIDInstructionsHandler == nil {
		api.GamePostGamesIDInstructionsHandler = game.PostGamesIDInstructionsHandlerFunc(func(params game.PostGamesIDInstructionsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation game.PostGamesIDInstructions has not yet been implemented")
		})
	}
	if api.ManagePostGamesIDJoinHandler == nil {
		api.ManagePostGamesIDJoinHandler = manage.PostGamesIDJoinHandlerFunc(func(params manage.PostGamesIDJoinParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation manage.PostGamesIDJoin has not yet been implemented")
		})
	}
	if api.ManagePostGamesNewHandler == nil {
		api.ManagePostGamesNewHandler = manage.PostGamesNewHandlerFunc(func(params manage.PostGamesNewParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation manage.PostGamesNew has not yet been implemented")
		})
	}
	if api.PostSignupCodeHandler == nil {
		api.PostSignupCodeHandler = operations.PostSignupCodeHandlerFunc(func(params operations.PostSignupCodeParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostSignupCode has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
