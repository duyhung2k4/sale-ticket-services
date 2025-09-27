package internal

import (
	"context"
	manager_api "sale-tickets/manager-service/gen"
)

type managerService struct {
	manager_api.UnimplementedHealthServiceServer
}

func (m *managerService) Ping(ctx context.Context, req *manager_api.Request) (*manager_api.Response, error) {
	return &manager_api.Response{
		Mess: "oke",
	}, nil
}

func NewManagerService() manager_api.HealthServiceServer {
	return &managerService{}
}
