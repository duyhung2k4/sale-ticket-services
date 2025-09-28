package main

import (
	"fmt"

	"sale-tickets/manager-service/internal/common/connection"
	"sale-tickets/manager-service/internal/router"
)

func main() {
	connection.Init()

	chanStartedGrpc := make(chan bool, 1)
	var chanError = make(chan error, 1)

	go func() {
		router.GrpcServer(chanStartedGrpc, chanError)
	}()

	go func() {
		router.HttpServer(chanStartedGrpc, chanError)
	}()

	err := <-chanError
	if err != nil {
		fmt.Println("error run main: ", err.Error())
		return
	}
}
