// Code generated by go-swagger; DO NOT EDIT.

package manage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostGamesIDJoinHandlerFunc turns a function with the right signature into a post games ID join handler
type PostGamesIDJoinHandlerFunc func(PostGamesIDJoinParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostGamesIDJoinHandlerFunc) Handle(params PostGamesIDJoinParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostGamesIDJoinHandler interface for that can handle valid post games ID join params
type PostGamesIDJoinHandler interface {
	Handle(PostGamesIDJoinParams, interface{}) middleware.Responder
}

// NewPostGamesIDJoin creates a new http.Handler for the post games ID join operation
func NewPostGamesIDJoin(ctx *middleware.Context, handler PostGamesIDJoinHandler) *PostGamesIDJoin {
	return &PostGamesIDJoin{Context: ctx, Handler: handler}
}

/* PostGamesIDJoin swagger:route POST /games/{id}/join manage postGamesIdJoin

Join an existing game

*/
type PostGamesIDJoin struct {
	Context *middleware.Context
	Handler PostGamesIDJoinHandler
}

func (o *PostGamesIDJoin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostGamesIDJoinParams()
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
