// Code generated by go-swagger; DO NOT EDIT.

package fizzbuzz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetAPIFizzbuzzOKCode is the HTTP code returned for type GetAPIFizzbuzzOK
const GetAPIFizzbuzzOKCode int = 200

/*GetAPIFizzbuzzOK fizz-buzz-like string based on given parameters

swagger:response getApiFizzbuzzOK
*/
type GetAPIFizzbuzzOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetAPIFizzbuzzOK creates GetAPIFizzbuzzOK with default headers values
func NewGetAPIFizzbuzzOK() *GetAPIFizzbuzzOK {

	return &GetAPIFizzbuzzOK{}
}

// WithPayload adds the payload to the get Api fizzbuzz o k response
func (o *GetAPIFizzbuzzOK) WithPayload(payload []string) *GetAPIFizzbuzzOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api fizzbuzz o k response
func (o *GetAPIFizzbuzzOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIFizzbuzzOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}