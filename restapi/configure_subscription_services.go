package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/kcraybould/subscriptions/restapi/operations"
	"github.com/kcraybould/subscriptions/restapi/operations/subscriptions"
)

// This file is safe to edit. Once it exists it will not be overwritten

func configureFlags(api *operations.SubscriptionServicesAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SubscriptionServicesAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.SubscriptionsDeleteHandler = subscriptions.DeleteHandlerFunc(func(params subscriptions.DeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation subscriptions.Delete has not yet been implemented")
	})
	api.SubscriptionsUnsubscribeHandler = subscriptions.UnsubscribeHandlerFunc(func(params subscriptions.UnsubscribeParams) middleware.Responder {
		return middleware.NotImplemented("operation subscriptions.Unsubscribe has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
