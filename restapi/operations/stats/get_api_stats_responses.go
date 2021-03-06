// Code generated by go-swagger; DO NOT EDIT.

package stats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Geoffrey42/fizzbuzz/models"
)

// GetAPIStatsOKCode is the HTTP code returned for type GetAPIStatsOK
const GetAPIStatsOKCode int = 200

/*GetAPIStatsOK A statistics endpoint allowing users to know what the most frequent request has been.

swagger:response getApiStatsOK
*/
type GetAPIStatsOK struct {

	/*
	  In: Body
	*/
	Payload *models.Stat `json:"body,omitempty"`
}

// NewGetAPIStatsOK creates GetAPIStatsOK with default headers values
func NewGetAPIStatsOK() *GetAPIStatsOK {

	return &GetAPIStatsOK{}
}

// WithPayload adds the payload to the get Api stats o k response
func (o *GetAPIStatsOK) WithPayload(payload *models.Stat) *GetAPIStatsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api stats o k response
func (o *GetAPIStatsOK) SetPayload(payload *models.Stat) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIStatsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAPIStatsNotFoundCode is the HTTP code returned for type GetAPIStatsNotFound
const GetAPIStatsNotFoundCode int = 404

/*GetAPIStatsNotFound No stored request can be found.

swagger:response getApiStatsNotFound
*/
type GetAPIStatsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAPIStatsNotFound creates GetAPIStatsNotFound with default headers values
func NewGetAPIStatsNotFound() *GetAPIStatsNotFound {

	return &GetAPIStatsNotFound{}
}

// WithPayload adds the payload to the get Api stats not found response
func (o *GetAPIStatsNotFound) WithPayload(payload *models.Error) *GetAPIStatsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api stats not found response
func (o *GetAPIStatsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIStatsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAPIStatsInternalServerErrorCode is the HTTP code returned for type GetAPIStatsInternalServerError
const GetAPIStatsInternalServerErrorCode int = 500

/*GetAPIStatsInternalServerError Database isn't available.

swagger:response getApiStatsInternalServerError
*/
type GetAPIStatsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAPIStatsInternalServerError creates GetAPIStatsInternalServerError with default headers values
func NewGetAPIStatsInternalServerError() *GetAPIStatsInternalServerError {

	return &GetAPIStatsInternalServerError{}
}

// WithPayload adds the payload to the get Api stats internal server error response
func (o *GetAPIStatsInternalServerError) WithPayload(payload *models.Error) *GetAPIStatsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api stats internal server error response
func (o *GetAPIStatsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIStatsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
