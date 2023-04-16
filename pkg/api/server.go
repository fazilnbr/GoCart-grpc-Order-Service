package http

import (
	"fmt"
	"log"
	"net"

	"github.com/fazilnbr/GoCart-grpc-order-Service/pkg/api/service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewGRPCServer(orderService *service.OrderService, grpcPort string) {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	fmt.Println("grpcPort/////", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// pb.RegisterAuthServiceServer(grpcServer, orderService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewServerHTTP(orderService *service.OrderService) *ServerHTTP {
	engine := gin.New()
	go NewGRPCServer(orderService, "50081")
	// Use logger from Gin
	engine.Use(gin.Logger())

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8000")
}
