package health

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/butterv/twirp-sample/app/gen/go/v1/health"
)

func (*healthService) Check(_ context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	if req.GetService() == "test" {
		return nil, status.Errorf(codes.InvalidArgument, "ng service")
	}

	// 	0: UNKNOWN
	//	1: SERVING
	//	2: NOT_SERVING
	//	3: SERVICE_UNKNOWN
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}, nil
}
