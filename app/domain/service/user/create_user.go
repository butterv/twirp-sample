package user

import (
	"context"

	pb "github.com/butterv/twirp-sample/app/gen/go/v1/user"
)

func (s *userService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	uID := s.userIDGenerator.Generate()

	return &pb.CreateUserResponse{
		UserId: string(uID),
	}, nil
}
