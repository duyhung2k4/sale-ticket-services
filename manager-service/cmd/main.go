package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	manager_api "sale-tickets/manager-service/gen"
	"sale-tickets/manager-service/internal"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func GrpcServer(startedGrpc chan<- bool, errStartGrpcServer chan<- error) {
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	manager_api.RegisterHealthServiceServer(s, internal.NewManagerService())

	log.Println("gRPC server running on :50050")
	startedGrpc <- true
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		errStartGrpcServer <- err
	}
}

func HttpServer(startedGrpc <-chan bool, errStartHttpServer chan<- error) {
	<-startedGrpc

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := manager_api.RegisterHealthServiceHandlerFromEndpoint(ctx, mux, "localhost:50050", opts)
	if err != nil {
		log.Fatalf("failed to register gateway: %v", err)
		errStartHttpServer <- err
	}

	log.Println("REST gateway running on :8000")
	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		errStartHttpServer <- err
	}
}

func main() {
	chanStartedGrpc := make(chan bool, 1)
	var chanError = make(chan error, 1)

	go func() {
		GrpcServer(chanStartedGrpc, chanError)
	}()

	go func() {
		HttpServer(chanStartedGrpc, chanError)
	}()

	err := <-chanError
	if err != nil {
		fmt.Println("error run main: ", err.Error())
		return
	}
}
