package api

import (
	"context"

	"github.com/microservices-spb/gateway/internal/model"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *model.User) (string, error)
	FindById(ctx context.Context, id int64) (*model.User, error)
}

type AuthClient interface {
	DoLogin(ctx context.Context, data model.RequestData) (string, error)
}
