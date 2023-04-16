package usecase

import (
	repository "github.com/fazilnbr/GoCart-grpc-order-Service/pkg/repository/interface"
	interfaces "github.com/fazilnbr/GoCart-grpc-order-Service/pkg/usecase/interface"
)

type orderUseCase struct {
	orderRepo repository.OrderRepository
}

func NewOrderUseCase(repo repository.OrderRepository) interfaces.OrderUseCase {
	return &orderUseCase{
		orderRepo: repo,
	}
}
