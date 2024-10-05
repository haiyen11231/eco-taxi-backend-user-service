# Project Structure

This project follows the [golang-standards/project-layout](https://github.com/golang-standards/project-layout) guidelines. Below is the structure of the project and an explanation of each folder and file.

## Directory Structure

```plaintext
eco-taxi-backend-user-service/
│
├── cmd/
│   └── user_service/
│       └── main.go               # Main entry point, initializes the service and starts the server
│
├── config/                       # Configuration settings (database, Redis, gRPC endpoints, etc.)
│   ├── config.go                 # Main config loading
│   ├── grpc_config.go            # gRPC-specific configuration
│   ├── mysql_config.go           # MySQL-specific configuration
│   └── redis.go                  # Redis-specific configuration
│
├── internal/
│   ├── auth/
│   │   ├── auth.go               # Handles authentication-related logic (JWT, token generation)
│   │   ├── middleware.go         # Middleware for authenticating incoming requests
│   │   └── tokens.go             # Token management (JWT access + refresh tokens)
│   │
│   ├── model/                    # Models represent the data structure (MVC: Model)
│   │   ├── user.go               # User entity definition (GORM model)
│   │   └── auth_response.go      # Auth response model (tokens, expiry)
│   │
│   ├── repository/               # Handles database operations (MVC: Model)
│   │   ├── user_repository.go    # User repository with GORM methods (CRUD operations)
│   │   └── token_repository.go   # Manages refresh token storage and retrieval
│   │
│   ├── service/                  # Business logic (MVC: Controller)
│   │   ├── user_service.go       # Core user-related logic (register, login, update profile)
│   │   ├── validation_service.go # Handles input validation (email, phone number, password)
│   │   ├── password_service.go   # Password hashing and validation
│   │   └── jwt_service.go        # Manages JWT token generation, validation, and expiry
│   │
│   ├── route/                    # Handles API endpoints (MVC: Controller)
│   │   ├── routes.go             # API route definitions (Gin handlers for each endpoint)
│   │   └── swagger.go            # Swagger configuration for API documentation
│   │
│   ├── grpc/                     # gRPC communication setup
│   │   ├── grpc_client.go        # gRPC client for communication with other services
│   │   ├── grpc_server.go        # gRPC server for handling incoming requests
│   │   └── proto/                # gRPC protobuf files
│   │       ├── user.proto        # .proto files defining gRPC methods for user service
│   │       └── other_protos.proto# Additional proto files
│   │
│   └── script/
│       └── migration/            # Database migration scripts (for versioning schema)
│
├── Dockerfile                    # Dockerfile for building the service container
├── docker-compose.yml            # Docker Compose file for local development with dependencies (optional if using Kubernetes)
├── Makefile                      # Build and run commands for the project
├── deployment/                   # Kubernetes manifests
│   ├── deployment.yaml           # Kubernetes deployment config for the service
│   ├── service.yaml              # Kubernetes service config (exposing the service)
│   └── ingress.yaml              # Ingress config for external access (optional)
│
├── docs/                         # API documentation
│   ├── api.md                    # API description and usage guide
│   ├── swagger.json              # Swagger JSON file for API documentation
│   └── swagger.yaml              # Swagger YAML file for API documentation
│
├── .gitignore                    # Git ignore file to exclude certain files from version control
└── README.md                     # Project documentation
```

## Directory Details

### `cmd/`

Contains the main application entry point.

- **`user_service/`**
  - **`main.go`**: Initializes and runs the User Service application.

### `config/`

Contains configuration settings for various services like the database, Redis, and gRPC.

- **`config.go`**: Main configuration loader, orchestrating the loading of individual configurations.
- **`grpc_config.go`**: gRPC-specific configuration settings.
- **`mysql_config.go`**: MySQL database configuration settings.
- **`redis.go`**: Redis-specific configuration settings.

### `internal/`

Contains the core application code, which includes authentication, business logic, database interactions, and gRPC setup.

- **`auth/`**

  - **`auth.go`**: Handles authentication-related logic, including JWT and token generation.
  - **`middleware.go`**: Middleware for authenticating incoming HTTP requests using JWT tokens.
  - **`tokens.go`**: Manages access and refresh token operations, including generation and validation.

- **`model/`**

  - **`user.go`**: Defines the User entity with GORM annotations.
  - **`auth_response.go`**: Represents the response structure for authentication (tokens, expiration, etc.).

- **`repository/`**

  - **`user_repository.go`**: Handles CRUD operations for user data using GORM.
  - **`token_repository.go`**: Manages storage and retrieval of refresh tokens in the database.

- **`service/`**

  - **`user_service.go`**: Core business logic related to user registration, login, profile management, etc.
  - **`validation_service.go`**: Validates inputs like email, phone number, and password.
  - **`password_service.go`**: Provides functionality for password hashing and validation.
  - **`jwt_service.go`**: Handles JWT token creation, validation, and expiration management.

- **`route/`**

  - **`routes.go`**: Defines all API routes, mapping each endpoint to the corresponding handler.
  - **`swagger.go`**: Configures Swagger for API documentation, providing UI access to explore and test the APIs.

- **`grpc/`**

  - **`grpc_client.go`**: Implements the gRPC client to communicate with other microservices.
  - **`grpc_server.go`**: Implements the gRPC server for handling incoming gRPC requests.
  - **`proto/`**
    - **`user.proto`**: Defines the gRPC methods and messages for the User Service.
    - **`other_protos.proto`**: Contains additional proto files for other related gRPC services.

- **`script/`**
  - **`migration/`**: Contains database migration scripts for versioning schema changes.

### `Dockerfile`

Dockerfile for building the User Service application image, specifying the base image and build steps.

### `docker-compose.yml`

Optional Docker Compose file for local development to manage dependencies like MySQL, Redis, etc., along with the user service.

### `Makefile`

Defines the build and run commands for the project, simplifying the development workflow.

### `deployment/`

Contains Kubernetes manifests for deploying the User Service in a Kubernetes cluster.

- **`deployment.yaml`**: Defines the deployment configuration, including replica settings and resource limits.
- **`service.yaml`**: Exposes the User Service within the Kubernetes cluster using a Kubernetes Service object.
- **`ingress.yaml`**: (Optional) Configures Ingress for external access to the service.

### `docs/`

Contains API documentation, including Swagger files for generating and viewing API specs.

- **`api.md`**: Detailed API description and usage guide.
- **`swagger.json`**: Swagger file in JSON format for API documentation.
- **`swagger.yaml`**: Swagger file in YAML format for API documentation.

### Root Directory

- **`.gitignore`**: Specifies which files and directories to exclude from version control (e.g., environment files, build artifacts).
- **`README.md`**: Project documentation and directory details.
- **`Dockerfile`**: Defines how to build the Docker image for the service.
- **`docker-compose.yml`**: Docker Compose file for managing the service and its dependencies.
- **`Makefile`**: Automates common tasks such as running, building, and testing the application.

### Setting Up

1. **Clone the Repository**

   Clone this repository to your local machine using:

   ```bash
   git clone https://github.com/haiyen11231/eco-taxi-backend-user-service.git
   cd eco-taxi-backend-user-service
   ```

### Create a `.env` File

Create a `.env` file in the root directory of the project. This file should contain the environment variables required for the application to run. Here's a sample `.env` file:

```env
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=yourdatabase
DB_ROOT_PASSWORD=yourrootpassword
REDIS_PASSWORD=yourredispassword
```

Update the values with your own configuration:

- **`DB_USER`**: MySQL username.
- **`DB_PASSWORD`**: MySQL user password.
- **`DB_NAME`**: MySQL database name.
- **`DB_ROOT_PASSWORD`**: MySQL root password.
- **`REDIS_PASSWORD`**: Redis password.

### Build and Run the Docker Application

Use Docker Compose to build and run the containers. This command will:

1. Build the Docker images for your application based on the `Dockerfile`.
2. Start the containers as defined in the `docker-compose.yml` file.

Run the following command in the root of your project directory:

```bash
docker-compose up --build
```

- `--build`: Forces Docker Compose to rebuild the images even if they are up-to-date.
- `--detach` or `-d`: Runs the containers in the background and prints the container IDs.
- `--remove-orphans`: Removes containers for services not defined in the `docker-compose.yml` file.

To build and run the Docker application, execute the following command:

```bash
docker-compose up --build
```

This command will build the Docker images as defined in the `Dockerfile` and `docker-compose.yml` file, then start the containers in the background. If you make changes to the Dockerfile or dependencies, you can re-run this command to rebuild the images and restart the containers.

After running the application, you can access the application at `http://localhost:8080`.

To stop and remove the application, networks, and volumes created by `docker-compose up`, use the following command:

```bash
docker-compose down
```

This command will stop the running application and remove them, along with the networks and volumes that were created. It is useful for cleaning up after development or when you want to ensure a fresh start.

If you need to stop the application without removing them, you can use:

```bash
docker-compose stop
```

This will stop the running application but leave them in place, so they can be restarted later.

To restart the application, you can use:

```bash
docker-compose start
```

This command starts the stopped application without rebuilding the images. It is a quick way to resume your application if you need to pause and resume development or testing.

To remove all application, networks, and volumes defined in your docker-compose.yml, you can use:

```bash
docker-compose down --volumes
```

This command stops and removes the application, networks, and volumes associated with it. It is useful for cleaning up your environment or ensuring a fresh start.

For more detailed information on Docker Compose commands and options, refer to the [official Docker Compose documentation](https://docs.docker.com/compose/).
