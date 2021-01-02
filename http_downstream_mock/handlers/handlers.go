package handlers

import (
	"net/http"
)

type handlers struct {
	client Client
}

type handlerOptions struct {
	client Client
}

// https://golang.org/pkg/net/http/#Client
// Here only embed the Get method
type Client interface {
	Get(string) (*http.Response, error)
}

// WithCustomerClient is a high-order function which sets the handler's client
func WithCustomerClient(client Client) func(*handlerOptions) *handlerOptions {
	return func(ho *handlerOptions) *handlerOptions {
		ho.client = client
		return ho
	}
}

// NewHandlers construct a handler with optional client. http.DefaultClient is
// used when no option is provided.
func NewHandlers(options ...func(*handlerOptions) *handlerOptions) *handlers {
	opts := &handlerOptions{
		client: http.DefaultClient,
	}
	for i := range options {
		opts = options[i](opts)
	}
	return &handlers{client: opts.client}
}
