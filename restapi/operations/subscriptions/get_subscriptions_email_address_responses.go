package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kcraybould/subscriptions/models"
)

/*GetSubscriptionsEmailAddressOK Success

swagger:response getSubscriptionsEmailAddressOK
*/
type GetSubscriptionsEmailAddressOK struct {

	// In: body
	Payload *models.Subscriber `json:"body,omitempty"`
}

// NewGetSubscriptionsEmailAddressOK creates GetSubscriptionsEmailAddressOK with default headers values
func NewGetSubscriptionsEmailAddressOK() *GetSubscriptionsEmailAddressOK {
	return &GetSubscriptionsEmailAddressOK{}
}

// WithPayload adds the payload to the get subscriptions email address o k response
func (o *GetSubscriptionsEmailAddressOK) WithPayload(payload *models.Subscriber) *GetSubscriptionsEmailAddressOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get subscriptions email address o k response
func (o *GetSubscriptionsEmailAddressOK) SetPayload(payload *models.Subscriber) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSubscriptionsEmailAddressOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetSubscriptionsEmailAddressBadRequest Bad Request

swagger:response getSubscriptionsEmailAddressBadRequest
*/
type GetSubscriptionsEmailAddressBadRequest struct {

	// In: body
	Payload *models.ErrorDetail `json:"body,omitempty"`
}

// NewGetSubscriptionsEmailAddressBadRequest creates GetSubscriptionsEmailAddressBadRequest with default headers values
func NewGetSubscriptionsEmailAddressBadRequest() *GetSubscriptionsEmailAddressBadRequest {
	return &GetSubscriptionsEmailAddressBadRequest{}
}

// WithPayload adds the payload to the get subscriptions email address bad request response
func (o *GetSubscriptionsEmailAddressBadRequest) WithPayload(payload *models.ErrorDetail) *GetSubscriptionsEmailAddressBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get subscriptions email address bad request response
func (o *GetSubscriptionsEmailAddressBadRequest) SetPayload(payload *models.ErrorDetail) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSubscriptionsEmailAddressBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetSubscriptionsEmailAddressNotFound Not Found

swagger:response getSubscriptionsEmailAddressNotFound
*/
type GetSubscriptionsEmailAddressNotFound struct {
}

// NewGetSubscriptionsEmailAddressNotFound creates GetSubscriptionsEmailAddressNotFound with default headers values
func NewGetSubscriptionsEmailAddressNotFound() *GetSubscriptionsEmailAddressNotFound {
	return &GetSubscriptionsEmailAddressNotFound{}
}

// WriteResponse to the client
func (o *GetSubscriptionsEmailAddressNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}