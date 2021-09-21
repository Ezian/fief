// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"fief/models"
)

// RegisterOKCode is the HTTP code returned for type RegisterOK
const RegisterOKCode int = 200

/*RegisterOK Successful registration

swagger:response registerOK
*/
type RegisterOK struct {

	/*
	  In: Body
	*/
	Payload *models.SuccessResponse `json:"body,omitempty"`
}

// NewRegisterOK creates RegisterOK with default headers values
func NewRegisterOK() *RegisterOK {

	return &RegisterOK{}
}

// WithPayload adds the payload to the register o k response
func (o *RegisterOK) WithPayload(payload *models.SuccessResponse) *RegisterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register o k response
func (o *RegisterOK) SetPayload(payload *models.SuccessResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterBadRequestCode is the HTTP code returned for type RegisterBadRequest
const RegisterBadRequestCode int = 400

/*RegisterBadRequest Bad Request

swagger:response registerBadRequest
*/
type RegisterBadRequest struct {
}

// NewRegisterBadRequest creates RegisterBadRequest with default headers values
func NewRegisterBadRequest() *RegisterBadRequest {

	return &RegisterBadRequest{}
}

// WriteResponse to the client
func (o *RegisterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// RegisterInternalServerErrorCode is the HTTP code returned for type RegisterInternalServerError
const RegisterInternalServerErrorCode int = 500

/*RegisterInternalServerError Server error

swagger:response registerInternalServerError
*/
type RegisterInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewRegisterInternalServerError creates RegisterInternalServerError with default headers values
func NewRegisterInternalServerError() *RegisterInternalServerError {

	return &RegisterInternalServerError{}
}

// WithPayload adds the payload to the register internal server error response
func (o *RegisterInternalServerError) WithPayload(payload string) *RegisterInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register internal server error response
func (o *RegisterInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
