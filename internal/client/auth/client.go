package auth

import (
	"context"
	"fmt"
	"log"

	"github.com/microservices-spb/auth/pkg/auth"
	"github.com/microservices-spb/gateway/internal/model"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client auth.AuthServiceClient
}

type CheckInDB struct {
	Check UserRepository
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

func (c *Client) SignUp(ctx context.Context, data model.RequestData) (string, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	username := model.User{
		Username: data.Username,
	}
	query := "SELECT id, username FROM userinfo WHERE username = $1"
	err := PostgresUserRepository.Conn.QueryRowContext(ctx, query, username).Scan(&username.Id, &username)

}
