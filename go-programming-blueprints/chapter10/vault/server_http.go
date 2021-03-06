package vault

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

// NewHTTPServer give a new http server
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	m := http.NewServeMux()
	// m.Handle("/hash", httptransport.NewServer(ctx, endpoints.HashEndpoint, decodeHashRequest, encodeResponse))
	m.Handle("/hash", httptransport.NewServer(endpoints.HashEndpoint, decodeHashRequest, encodeResponse))
	// m.Handle("/validate", httptransport.NewServer(ctx, endpoints.ValidateEndpoint, decodeValidateRequest, encodeResponse))
	m.Handle("/validate", httptransport.NewServer(endpoints.ValidateEndpoint, decodeValidateRequest, encodeResponse))
	return m
}
