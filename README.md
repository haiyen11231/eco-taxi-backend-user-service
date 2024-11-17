# EcoTaxi Backend - User Service

EcoTaxi Backend - User Service is responsible for handling user-related requests received from the API Gateway. It processes these requests using [Gin](https://github.com/gin-gonic/gin) for routing, [GORM](https://gorm.io/docs/) for interacting with [MySQL](https://dev.mysql.com/doc/), and [Redis](https://redis.io/docs/latest/develop/get-started/) to cache user tokens.

## Git Repositories

This project is part of the EcoTaxi ecosystem, which includes multiple repositories for the frontend, backend services, and API gateway:

- **Frontend**: [EcoTaxi Frontend](https://github.com/haiyen11231/eco-taxi-frontend.git)
- **API Gateway**: [EcoTaxi API Gateway](https://github.com/haiyen11231/eco-taxi-api-gateway.git)
- **User Service**: [EcoTaxi User Service](https://github.com/haiyen11231/eco-taxi-backend-user-service.git)
- **Payment Service**: [EcoTaxi Payment Service](https://github.com/AWYS7/eco-taxi-payment-service.git)
- **Trip Service**: [EcoTaxi Trip Service](https://github.com/lukea11/eco-taxi-backend-trip-service.git)

## Directory Structure

```plaintext
eco-taxi-backend-user-service/
│
├── cmd/
│   └── user_service/
│       └── main.go
│
├── config/
│   ├── config.go
│   ├── grpc_config.go
│   ├── mysql_config.go
│   └── redis_config.go
│
├── internal/
│   ├── cache/
│   │   └── session_cache.go
│   │
│   ├── model/
│   │   └── user.go
│   │
│   ├── repository/
│   │   └── user_repository.go
│   │
│   ├── service/
│   │   ├── jwt_service.go
│   │   └── user_service.go
│   │
│   ├── route/
│   │   ├── routes.go
│   │   └── swagger.go
│   │
│   ├── grpc/
│   │   ├── user_service.proto
│   │   └── pb/
│   │       ├── user_service_grpc.pb.go
│   │       └── user_service.pb.go
│   │
│   ├── script/
│   │   └── migrations/
│   │
│   └── utils/
│       └── email.go
│
├── deployment/
│   ├── deployment.yaml
│   ├── service.yaml
│   └── ingress.yaml
│
├── docs/
│   ├── api.md
│   ├── swagger.json
│   └── swagger.yaml
│
├── .gitignore
├── app.env
├── Dockerfile
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Prerequisites

Before you begin, ensure that you have the following installed:

- **Go**
- **gRPC Tools** (Protocol Buffers and gRPC Go)
- **MySQL**
- **Redis**
- **Make**
- **Docker** (optional, for containerization)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/haiyen11231/eco-taxi-backend-user-service.git
   cd eco-taxi-backend-user-service
   ```

2. Create the app.env file:

Create a `app.env` file in the root directory of the project. This file should contain the environment variables required for the application to run. Here's a sample `app.env` file:

```env
# Database configuration
MYSQL_HOST=mysql_host
MYSQL_PORT=mysql_port
MYSQL_USER=mysql_user
MYSQL_PASSWORD=mysql_password
MYSQL_DATABASE=mysql_db

# Redis configuration
REDIS_HOST=redis_host
REDIS_PORT=redis_port
REDIS_PASSWORD=redis_password
REDIS_DB=0

# gRPC configuration
GRPC_PORT=grpc_port

JWT_SECRET=secret

PORT=port
```

Update the values with your own configuration:

- **`MYSQL_*`**: MySQL configuration (host, port, user, password, and database).
- **`REDIS_*`**: Redis configuration (host, port, password, and DB number).
- **`GRPC_PORT`**: Port on which the gRPC server for User Service will run (e.g., localhost:5002).
- **`JWT_SECRET`**: Secret key used for signing and verifying JWT tokens.
- **`PORT`**: Define the port number on which the User Service API will listen (e.g., 8082).

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Start the development server:

   ```bash
   make run
   ```
