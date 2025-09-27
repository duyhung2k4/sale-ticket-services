package health_controller

import (
	"context"
	manager_api "sale-tickets/manager-service/gen"
)

func (m *managerService) Ping(ctx context.Context, req *manager_api.Request) (*manager_api.Response, error) {
	return &manager_api.Response{
		Mess: "oke",
	}, nil
}
