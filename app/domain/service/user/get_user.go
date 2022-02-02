package user

import (
	"context"
	"time"

	"github.com/butterv/twirp-sample/app/domain/model"
	pb "github.com/butterv/twirp-sample/app/gen/go/v1/user"
	"github.com/butterv/twirp-sample/app/presenter"
)

func (s *userService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	uID := model.UserID(req.GetUserId())

	u := &model.User{
		ID:        uID,
		Email:     "test@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	return &pb.GetUserResponse{
		User: presenter.UserToPbUser(u),
	}, nil
}
