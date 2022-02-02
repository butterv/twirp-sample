package main

import (
	"net/http"

	"github.com/butterv/twirp-sample/app/domain/model"
	"github.com/butterv/twirp-sample/app/domain/service/health"
	"github.com/butterv/twirp-sample/app/domain/service/user"
	healthpb "github.com/butterv/twirp-sample/app/gen/go/v1/health"
	userpb "github.com/butterv/twirp-sample/app/gen/go/v1/user"
)

func main() {
	mux := http.NewServeMux()

	healthTwirp := healthpb.NewHealthServer(health.NewHealthService())
	userTwirp := userpb.NewUsersServer(user.NewUserService(model.NewDefaultUserIDGenerator()))

	mux.Handle(healthTwirp.PathPrefix(), healthTwirp)
	mux.Handle(userTwirp.PathPrefix(), userTwirp)

	http.ListenAndServe(":8080", mux)
}
