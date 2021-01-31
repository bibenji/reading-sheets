package vault

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"golang.org/x/net/context"

	pb "../vault/pb"
)

type grpcServer struct {
	hash     grpctransport.Handler
	validate grpctransport.Handler
}

// NewGRPCServer gets a new pb.VaultServer.
func NewGRPCServer(ctx context.Context, endpoints Endpoints) pb.VaultServer {
	return &grpcServer{
		hash: grpctransport.NewServer(
			ctx,
			endpoints.HashEndpoint,
			DecodeGRPCHashRequest,
			EncodeGRPCHashResponse,
		),
		validate: grpctransport.NewServer(
			ctx,
			endpoints.ValidateEndpoint,
			DecodeGRPCValidateRequest,
			EncodeGRPCValidateResponse,
		),
	}
}

func (s *grpcServer) Hash(ctx context.Context, r *pb.HashRequest) (*pb.HashResponse, error) {
	_, resp, err := s.hash.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.HashResponse), nil
}

func (s *grpctransport) Validate(ctx context.Context, r *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	_, resp, err := s.validate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ValidateResponse), nil
}

// EncodeGRPCHashRequest todo
func EncodeGRPCHashRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(hashRequest)
	return &pb.HashRequest{Password: req.Password}, nil
}

// DecodeGRPCHashRequest todo
func DecodeGRPCHashRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.HashRequest)
	return hashRequest{Password: req.Password}, nil
}

// EncodeGRPCHashResponse todo
func EncodeGRPCHashResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(hashResponse)
	return &pb.HashResponse{Hash: res.Hash, Err: res.Err}, nil
}

// DecodeGRPCHashResponse todo
func DecodeGRPCHashResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.HashResponse)
	return hashResponse{Hash: res.Hash, Err: res.Err}, nil
}

// EncodeGRPCValidateRequest todo
func EncodeGRPCValidateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(validateRequest)
	return &pb.ValidateRequest{Password: req.Password, Hash: req.Hash}, nil
}

// DecodeGRPCValidateRequest todo
func DecodeGRPCValidateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.ValidateRequest)
	return validateRequest{Password: req.Password, Hash: req.Hash}, nil
}

// EncodeGRPCValidateResponse todo
func EncodeGRPCValidateResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(validateResponse)
	return &pb.ValidateResponse{Valid: res.Valid}, nil
}

// DecodeGRPCValidateResponse todo
func DecodeGRPCValidateResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.ValidateResponse)
	return validateResponse{Valid: res.Valid}, nil
}
