package main

import (
	"fmt"
	"folder-structure/internal/api"
	"folder-structure/internal/repository"
	"folder-structure/internal/service"
	"log"
	"net/http"
)

func main() {

	repo := repository.New()

	srv := service.New(repo)

	handler := api.New(srv)

	response := handler.Do(2, 7)

	fmt.Println(response)

	http.HandleFunc("/{q}", handler.Handler)

	err := http.ListenAndServe(":3111", nil)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
