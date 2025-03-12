package main

import (
	"log"
	"net/http"

	"github.com/microservices-spb/gateway/internal/api"
	"github.com/microservices-spb/gateway/internal/client/auth"
	"github.com/microservices-spb/gateway/internal/repository"
)

func main() {

	conn := repository.ConnectToDB()

	userRepo := repository.NewPostgresUserRepository(conn.Conn)

	authClient := auth.New()

	handler := api.New(userRepo, authClient)

	http.HandleFunc("/login", handler.Handler)

	log.Println("service start...")

	err := http.ListenAndServe(":3111", nil)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
