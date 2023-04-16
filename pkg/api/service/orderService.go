package service

import (
	usecase "github.com/fazilnbr/GoCart-grpc-order-Service/pkg/usecase/interface"
)

type OrderService struct {
	orderUseCase usecase.OrderUseCase
}

func NewOrderService(usecase usecase.OrderUseCase) *OrderService {
	return &OrderService{
		orderUseCase: usecase,
	}
}
