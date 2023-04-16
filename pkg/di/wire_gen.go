// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/fazilnbr/GoCart-grpc-order-Service/pkg/api"
	"github.com/fazilnbr/GoCart-grpc-order-Service/pkg/api/service"
	"github.com/fazilnbr/GoCart-grpc-order-Service/pkg/config"
	"github.com/fazilnbr/GoCart-grpc-order-Service/pkg/db"
	"github.com/fazilnbr/GoCart-grpc-order-Service/pkg/repository"
	"github.com/fazilnbr/GoCart-grpc-order-Service/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	orderRepository := repository.NewOrderRepository(gormDB)
	orderUseCase := usecase.NewOrderUseCase(orderRepository)
	orderService := service.NewOrderService(orderUseCase)
	serverHTTP := http.NewServerHTTP(orderService)
	return serverHTTP, nil
}