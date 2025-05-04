package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/microservices-spb/auth/pkg/auth"
	"github.com/microservices-spb/gateway/internal/api"
	"github.com/microservices-spb/gateway/internal/model"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client auth.AuthServiceClient
	check  api.UserRepository
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

func (c *Client) SignUp(ctx context.Context, data model.RequestData) (bool, error) {
	username, err := c.check.CheckUserInDB(ctx, data.Username)
	if err != nil {
		return username, err
	}
	if username {
		return username, errors.New("username already taken")
	}

	usernameStr := strconv.FormatBool(username)

	passHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("failed to generate password")
	}

	c.check.SaveUser(ctx, &model.RequestData{
		Username: usernameStr,
		Password: string(passHash),
	})

	return username, nil
}
