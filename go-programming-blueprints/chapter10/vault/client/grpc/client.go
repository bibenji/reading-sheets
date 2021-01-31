package grpc

import (
	vault "../../"
	pb "../../pb"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	"google.golang.org/grpc"
)

// New create a new
func New(conn *grpc.ClientConn) vault.Service {
	var hashEndpoint = grpctransport.NewClient(
		conn,
		"Vault",
		"Hash",
		vault.EncodeGRPCHashRequest,
		vault.DecodeGRPCHashResponse,
		pb.HashResponse{},
	).Endpoint()

	var validateEndpoint = grpctransport.NewClient(
		conn,
		"Vault",
		"Validate",
		vault.EncodeGRPCValidateRequest,
		vault.DecodeGRPCValidateResponse,
		pb.ValidateResponse{},
	).Endpoint()

	return vault.Endpoints{
		hashEndpoint:     hashEndpoint,
		validateEndpoint: validateEndpoint,
	}
}
