package main

import (
	"log"
	"net/http"

	"github.com/microservices-spb/gateway/internal/api"
	"github.com/microservices-spb/gateway/internal/client/auth"
	"github.com/microservices-spb/gateway/internal/repository"
	"github.com/microservices-spb/gateway/internal/service"
)

func main() {

	repo := repository.New()

	authClient := auth.New()

	srv := service.New(repo)

	handler := api.New(srv, authClient)

	//response := handler.Do(2, 7)

	//fmt.Println(response)

	http.HandleFunc("/login", handler.Handler)

	log.Println("service start...")

	err := http.ListenAndServe(":3111", nil)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
