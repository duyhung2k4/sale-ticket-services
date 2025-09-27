package router

import (
	"fmt"
	"log"
	"net"
	manager_api "sale-tickets/manager-service/gen"
	"sale-tickets/manager-service/internal"
	"sale-tickets/manager-service/internal/connection"

	"google.golang.org/grpc"
)

func GrpcServer(startedGrpc chan<- bool, errStartGrpcServer chan<- error) {
	port := fmt.Sprintf(":%s", connection.ConfigInfo.App.GrpcPort)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	manager_api.RegisterHealthServiceServer(s, internal.NewManagerService())

	log.Printf("gRPC server running on %s", port)
	startedGrpc <- true
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		errStartGrpcServer <- err
	}
}
