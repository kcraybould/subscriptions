package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutSubscriptionsEmailAddressHandlerFunc turns a function with the right signature into a put subscriptions email address handler
type PutSubscriptionsEmailAddressHandlerFunc func(PutSubscriptionsEmailAddressParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutSubscriptionsEmailAddressHandlerFunc) Handle(params PutSubscriptionsEmailAddressParams) middleware.Responder {
	return fn(params)
}

// PutSubscriptionsEmailAddressHandler interface for that can handle valid put subscriptions email address params
type PutSubscriptionsEmailAddressHandler interface {
	Handle(PutSubscriptionsEmailAddressParams) middleware.Responder
}

// NewPutSubscriptionsEmailAddress creates a new http.Handler for the put subscriptions email address operation
func NewPutSubscriptionsEmailAddress(ctx *middleware.Context, handler PutSubscriptionsEmailAddressHandler) *PutSubscriptionsEmailAddress {
	return &PutSubscriptionsEmailAddress{Context: ctx, Handler: handler}
}

/*PutSubscriptionsEmailAddress swagger:route PUT /subscriptions/{emailAddress} subscriptions putSubscriptionsEmailAddress

Update an email subscriber's profile and subscription information.  This is a composite service.


*/
type PutSubscriptionsEmailAddress struct {
	Context *middleware.Context
	Handler PutSubscriptionsEmailAddressHandler
}

func (o *PutSubscriptionsEmailAddress) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPutSubscriptionsEmailAddressParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
