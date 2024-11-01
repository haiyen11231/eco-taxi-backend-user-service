package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/haiyen11231/eco-taxi-backend-user-service/config"
	"github.com/haiyen11231/eco-taxi-backend-user-service/internal/grpc/pb"
	"github.com/haiyen11231/eco-taxi-backend-user-service/internal/service"
	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	r := gin.Default()
	log.Println("Starting User Service")
	
	
	if err := config.ConnectToMySQL(); err != nil {
		log.Panic("Failed to connect to MySQL:", err)
	}

	
	if err := config.ConnectToRedis(); err != nil {
		log.Panic("Failed to connect to Redis:", err)
	}

	go listenGRPC()

	
	if err := r.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		log.Panic(err)
	}
	
	// r.Run() // listen and serve on 0.0.0.0:8082 (for windows "localhost:8082")
}

func loadEnv() {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func listenGRPC() {
	gRPCPort := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC Auth: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, &service.UserServiceServer{})

	log.Printf("gRPC Auth Server started on port %s", gRPCPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC Auth: %v", err)
	}
}