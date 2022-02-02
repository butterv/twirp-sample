package user

import (
	"github.com/butterv/twirp-sample/app/domain/model"
	pb "github.com/butterv/twirp-sample/app/gen/go/v1/user"
)

type userService struct {
	userIDGenerator model.UserIDGenerator
}

// NewUserService generates the `UsersServer` implementation.
func NewUserService(userIDGenerator model.UserIDGenerator) pb.UsersServer {
	return &userService{
		userIDGenerator: userIDGenerator,
	}
}
