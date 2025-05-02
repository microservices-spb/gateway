package api

import (
	"context"

	"github.com/microservices-spb/gateway/internal/model"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *model.RequestData) (string, error)
	FindById(ctx context.Context, id int64) (*model.User, error)
	CheckUserInDB(ctx context.Context, username string) (bool, error)
}

type AuthClient interface {
	DoLogin(ctx context.Context, data model.RequestData) (string, error)
	SignUp(ctx context.Context, data model.RequestData) error
}
