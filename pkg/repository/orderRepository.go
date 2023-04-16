package repository

import (
	interfaces "github.com/fazilnbr/GoCart-grpc-order-Service/pkg/repository/interface"
	"gorm.io/gorm"
)

type orderDatabase struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepository {
	return &orderDatabase{DB}
}
