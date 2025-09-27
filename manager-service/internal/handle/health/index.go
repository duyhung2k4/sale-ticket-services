package health_controller

import (
	manager_api "sale-tickets/manager-service/gen"
)

type managerService struct {
	manager_api.UnimplementedHealthServer
}

func NewHandle() manager_api.HealthServer {
	return &managerService{}
}
