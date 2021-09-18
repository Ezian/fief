// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"fief/models"
)

// PostSignupCodeOKCode is the HTTP code returned for type PostSignupCodeOK
const PostSignupCodeOKCode int = 200

/*PostSignupCodeOK Successful login

swagger:response postSignupCodeOK
*/
type PostSignupCodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.LoginSuccess `json:"body,omitempty"`
}

// NewPostSignupCodeOK creates PostSignupCodeOK with default headers values
func NewPostSignupCodeOK() *PostSignupCodeOK {

	return &PostSignupCodeOK{}
}

// WithPayload adds the payload to the post signup code o k response
func (o *PostSignupCodeOK) WithPayload(payload *models.LoginSuccess) *PostSignupCodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post signup code o k response
func (o *PostSignupCodeOK) SetPayload(payload *models.LoginSuccess) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSignupCodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostSignupCodeBadRequestCode is the HTTP code returned for type PostSignupCodeBadRequest
const PostSignupCodeBadRequestCode int = 400

/*PostSignupCodeBadRequest Bad Request

swagger:response postSignupCodeBadRequest
*/
type PostSignupCodeBadRequest struct {
}

// NewPostSignupCodeBadRequest creates PostSignupCodeBadRequest with default headers values
func NewPostSignupCodeBadRequest() *PostSignupCodeBadRequest {

	return &PostSignupCodeBadRequest{}
}

// WriteResponse to the client
func (o *PostSignupCodeBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PostSignupCodeNotFoundCode is the HTTP code returned for type PostSignupCodeNotFound
const PostSignupCodeNotFoundCode int = 404

/*PostSignupCodeNotFound User not found

swagger:response postSignupCodeNotFound
*/
type PostSignupCodeNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostSignupCodeNotFound creates PostSignupCodeNotFound with default headers values
func NewPostSignupCodeNotFound() *PostSignupCodeNotFound {

	return &PostSignupCodeNotFound{}
}

// WithPayload adds the payload to the post signup code not found response
func (o *PostSignupCodeNotFound) WithPayload(payload string) *PostSignupCodeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post signup code not found response
func (o *PostSignupCodeNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSignupCodeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostSignupCodeInternalServerErrorCode is the HTTP code returned for type PostSignupCodeInternalServerError
const PostSignupCodeInternalServerErrorCode int = 500

/*PostSignupCodeInternalServerError Server error

swagger:response postSignupCodeInternalServerError
*/
type PostSignupCodeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostSignupCodeInternalServerError creates PostSignupCodeInternalServerError with default headers values
func NewPostSignupCodeInternalServerError() *PostSignupCodeInternalServerError {

	return &PostSignupCodeInternalServerError{}
}

// WithPayload adds the payload to the post signup code internal server error response
func (o *PostSignupCodeInternalServerError) WithPayload(payload string) *PostSignupCodeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post signup code internal server error response
func (o *PostSignupCodeInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSignupCodeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
