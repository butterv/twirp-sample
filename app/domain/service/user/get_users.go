package user

import (
	"context"
	"time"

	"github.com/butterv/twirp-sample/app/domain/model"
	pb "github.com/butterv/twirp-sample/app/gen/go/v1/user"
	"github.com/butterv/twirp-sample/app/presenter"
)

func (s *userService) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	var uIDs []model.UserID
	for _, uID := range req.GetUserIds() {
		uIDs = append(uIDs, model.UserID(uID))
	}

	us := make([]*model.User, len(uIDs))
	for _, uID := range uIDs {
		us = append(us, &model.User{
			ID:        uID,
			Email:     "test@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		})
	}

	return &pb.GetUsersResponse{
		Users: presenter.UsersToPbUsers(us),
	}, nil
}
