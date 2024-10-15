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
	
	err := config.ConnectToMySQL()
	if err != nil {
		log.Panic("Failed to connect to MySQL:", err)
	}

	go listenGRPC()

	err = r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
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

// jsonData, err := json.Marshal(user)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(jsonData))

	// jsonStr := "{\"id\":2,\"name\":\"Rose\",\"phone_number\":83641890,\"email\":\"rose@gmail.com\",\"hashed_password\":\"janeK9*bce\",\"distance_travelled\":7.5}"

	// var user2 User
	// if err := json.Unmarshal([]byte(jsonStr), &user2); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(user2)