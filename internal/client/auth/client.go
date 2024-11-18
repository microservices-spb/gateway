package auth

import (
	"context"
	"fmt"
	"github.com/microservices-spb/auth/pkg/auth"
	"github.com/microservices-spb/gateway/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Client struct {
	client auth.AuthServiceClient
}

func New() *Client {
	conn, err := grpc.NewClient("localhost:3112", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := auth.NewAuthServiceClient(conn)
	return &Client{client: client}
}

func (c *Client) DoLogin(ctx context.Context, data model.RequestData) (string, error) {
	resp, err := c.client.Login(ctx, &auth.LoginIn{
		Username: data.Username,
		Password: data.Password,
	})
	if err != nil {
		return "", fmt.Errorf("failed to login: %w", err)
	}
	return resp.Token, nil
}
