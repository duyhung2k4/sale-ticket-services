package main

import (
	"fmt"
	"log"
	"net/http"
	"sale-ticket-iam-service/internal/common/connection"
	"sale-ticket-iam-service/internal/router"
	"time"
)

func main() {
	connection.Init()

	config := connection.GetConfig()

	server := http.Server{
		Addr:           fmt.Sprintf("%s:%s", config.App.Host, config.App.Port),
		Handler:        router.AppRouter(),
		ReadTimeout:    600 * time.Second,
		WriteTimeout:   600 * time.Second,
		MaxHeaderBytes: 5 * 1024 * 1024,
	}

	log.Printf("Server listen: %s:%s", config.App.Host, config.App.Port)
	log.Fatalln(server.ListenAndServe())
}
