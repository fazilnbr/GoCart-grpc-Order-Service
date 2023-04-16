//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/fazilnbr/GoCart-grpc-order-Service/pkg/api"
	"github.com/fazilnbr/GoCart-grpc-order-Service/pkg/api/service"
	"github.com/fazilnbr/GoCart-grpc-order-Service/pkg/config"
	"github.com/fazilnbr/GoCart-grpc-order-Service/pkg/db"
	"github.com/fazilnbr/GoCart-grpc-order-Service/pkg/repository"
	"github.com/fazilnbr/GoCart-grpc-order-Service/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewOrderRepository,
		usecase.NewOrderUseCase,
		service.NewOrderService,
		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
