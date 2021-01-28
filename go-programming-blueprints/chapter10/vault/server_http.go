package vault

import (
	"net/http"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

// NewHTTPServer give a new http server
funct NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	m := htttp.NewServeMux()
	m.Handle("/hash", httptransport.NewServer(ctx, endpoints.HashEndpoint, decodeHashRequest, encodeResponse))
	m.Handle("/validate", httptransport.NewServer(ctx, endpoints.ValidateEndpoint, decodeValidateRequest, encodeResponse))
	return m
}
