// Code generated by go-swagger; DO NOT EDIT.

package game

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostGamesIDInstructionsHandlerFunc turns a function with the right signature into a post games ID instructions handler
type PostGamesIDInstructionsHandlerFunc func(PostGamesIDInstructionsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostGamesIDInstructionsHandlerFunc) Handle(params PostGamesIDInstructionsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostGamesIDInstructionsHandler interface for that can handle valid post games ID instructions params
type PostGamesIDInstructionsHandler interface {
	Handle(PostGamesIDInstructionsParams, interface{}) middleware.Responder
}

// NewPostGamesIDInstructions creates a new http.Handler for the post games ID instructions operation
func NewPostGamesIDInstructions(ctx *middleware.Context, handler PostGamesIDInstructionsHandler) *PostGamesIDInstructions {
	return &PostGamesIDInstructions{Context: ctx, Handler: handler}
}

/* PostGamesIDInstructions swagger:route POST /games/{id}/instructions game postGamesIdInstructions

Post an instruction for the authenticated user on the session, if no instruction defined

*/
type PostGamesIDInstructions struct {
	Context *middleware.Context
	Handler PostGamesIDInstructionsHandler
}

func (o *PostGamesIDInstructions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostGamesIDInstructionsParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
