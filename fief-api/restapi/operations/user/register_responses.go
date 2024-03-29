// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// RegisterOKCode is the HTTP code returned for type RegisterOK
const RegisterOKCode int = 200

/*RegisterOK Successful registration

swagger:response registerOK
*/
type RegisterOK struct {
}

// NewRegisterOK creates RegisterOK with default headers values
func NewRegisterOK() *RegisterOK {

	return &RegisterOK{}
}

// WriteResponse to the client
func (o *RegisterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// RegisterBadRequestCode is the HTTP code returned for type RegisterBadRequest
const RegisterBadRequestCode int = 400

/*RegisterBadRequest Registration failure

swagger:response registerBadRequest
*/
type RegisterBadRequest struct {

	/*Error message
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewRegisterBadRequest creates RegisterBadRequest with default headers values
func NewRegisterBadRequest() *RegisterBadRequest {

	return &RegisterBadRequest{}
}

// WithPayload adds the payload to the register bad request response
func (o *RegisterBadRequest) WithPayload(payload string) *RegisterBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register bad request response
func (o *RegisterBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// RegisterInternalServerErrorCode is the HTTP code returned for type RegisterInternalServerError
const RegisterInternalServerErrorCode int = 500

/*RegisterInternalServerError Server error

swagger:response registerInternalServerError
*/
type RegisterInternalServerError struct {

	/*Error message
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
