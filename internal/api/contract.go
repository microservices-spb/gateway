package api

import (
	"context"
	"github.com/microservices-spb/gateway/internal/model"
)

type Srv interface {
	Mulity(x, y int64) int64
}

type AuthClient interface {
	DoLogin(ctx context.Context, data model.RequestData) (string, error)
}
