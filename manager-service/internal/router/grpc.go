package router

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	manager_api "sale-tickets/manager-service/gen"

	"sale-tickets/manager-service/internal/common/connection"
	health_controller "sale-tickets/manager-service/internal/handle/health"
	moviethreater_controller "sale-tickets/manager-service/internal/handle/movie_threater"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func GrpcServer(startedGrpc chan<- bool, errStartGrpcServer chan<- error) {
	port := fmt.Sprintf(":%s", connection.ConfigInfo.App.GrpcPort)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
				md, ok := metadata.FromIncomingContext(ctx)
				if !ok {
					return nil, errors.New("metadata not ok")
				}

				md.Set("uuid", "0")
				ctx = metadata.NewIncomingContext(ctx, md)
				return handler(ctx, req)
			},
		),
	)
	manager_api.RegisterHealthServer(s, health_controller.NewHandle())
	manager_api.RegisterMovieTheaterServer(s, moviethreater_controller.NewHandle())

	log.Printf("gRPC server running on %s", port)
	startedGrpc <- true
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		errStartGrpcServer <- err
	}
}
