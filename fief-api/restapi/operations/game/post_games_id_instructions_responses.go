// Code generated by go-swagger; DO NOT EDIT.

package game

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostGamesIDInstructionsOKCode is the HTTP code returned for type PostGamesIDInstructionsOK
const PostGamesIDInstructionsOKCode int = 200

/*PostGamesIDInstructionsOK TODO

swagger:response postGamesIdInstructionsOK
*/
type PostGamesIDInstructionsOK struct {
}

// NewPostGamesIDInstructionsOK creates PostGamesIDInstructionsOK with default headers values
func NewPostGamesIDInstructionsOK() *PostGamesIDInstructionsOK {

	return &PostGamesIDInstructionsOK{}
}

// WriteResponse to the client
func (o *PostGamesIDInstructionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
