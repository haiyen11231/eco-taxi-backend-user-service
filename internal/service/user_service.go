package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/haiyen11231/eco-taxi-backend-user-service/config"
	"github.com/haiyen11231/eco-taxi-backend-user-service/internal/grpc/pb"
	"github.com/haiyen11231/eco-taxi-backend-user-service/internal/model"
	"github.com/haiyen11231/eco-taxi-backend-user-service/internal/repository"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServiceServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	if req.Name == "" || req.PhoneNumber == "" || req.Email == "" || req.Password == "" {
		return nil, errors.New("Name, Phone Number, Email and Password are required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password:", err.Error())
		return nil, err
	}

	signUpData := &model.SignUpUserData{
		Name:      req.Name,
		PhoneNumber: req.PhoneNumber,
		Email: req.Email,
		Password:  string(hashedPassword),
	}

	db := config.DB
	userRepo := repository.NewUserRepo(db)

	if err := userRepo.SignUp(ctx, signUpData); err != nil {
		log.Println("Failed to signup:", err.Error())
		return nil, err
	}

	return &pb.SignUpResponse{Message: "User created!"}, nil
}
	
func (s *UserServiceServer) LogIn(ctx context.Context, req *pb.LogInRequest) (*pb.LogInResponse, error) {
	if req.PhoneNumber == "" || req.Password == "" {
		return nil, errors.New("Phone Number and Password are required")
	}

	logInData := &model.LogInUserData{
		PhoneNumber: req.PhoneNumber,
		Password:  string(req.Password),
	}

	db := config.DB
	userRepo := repository.NewUserRepo(db)

	user, err := userRepo.LogIn(ctx, logInData)

	if err != nil {
		log.Println("Failed to login:", err.Error())
		return nil, err
	}

	token, err := generateToken(user)
	if err != nil {
		log.Println("Failed to generate token:", err.Error())
		return nil, errors.New("Invalid Phone Number or Password")
	}

	return &pb.LogInResponse{Token: token}, nil
}

func (s *UserServiceServer) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {
	if req.Email == "" || req.NewPassword == "" {
		return nil, errors.New("Missing required fields")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password:", err.Error())
		return nil, err
	}

	forgotPasswordUserData := &model.ChangePasswordUserData{
		NewPassword: string(hashedPassword),
	}

	db := config.DB
	userRepo := repository.NewUserRepo(db)

	if err := userRepo.ForgotPassword(ctx, forgotPasswordUserData, req.Email); err != nil {
		log.Println("Failed to reset password:", err.Error())
		return nil, err
	}

	return &pb.ForgotPasswordResponse{Message: "Password resetted!"}, nil
}

func (s *UserServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	if req.Id == 0 || req.Name == "" || req.PhoneNumber == "" || req.Email == "" {
		return nil, errors.New("Missing required fields")
	}

	updateData := &model.UpdateUserData{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Email: req.Email,
	}

	db := config.DB
	userRepo := repository.NewUserRepo(db)

	if err := userRepo.UpdateUser(ctx, updateData, uint64(req.Id)); err != nil {
		log.Println("Failed to update user:", err.Error())
		return nil, err
	}

	return &pb.UpdateUserResponse{Message: "User updated!"}, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var user model.User;
	// Checking if the request contains a valid user ID
	if req.Id == 0 {
		return nil, errors.New("Id is required")
	}

	user.Id = req.Id

	// Establishing a database connection and initializing the repository
	db := config.DB
	userRepo := repository.NewUserRepo(db)

	// Retrieving the user details from the database
	err := userRepo.GetUser(ctx, &user)

	if err != nil {
		log.Println("Failed to get user:", err.Error())
		return nil, err
	}

	// Mapping model.User to pb.GetUserResponse
	log.Println("User's Distance Travelled:", user.DistanceTravelled)
	userResponse := &pb.GetUserResponse{
		Id: user.Id,
		Name: user.Name,
		PhoneNumber: user.PhoneNumber,
		Email: user.Email,
		DistanceTravelled: user.DistanceTravelled,
	}

	return userResponse, nil
	
}

func (s *UserServiceServer) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	if req.Id ==0 || req.OldPassword == "" || req.NewPassword == "" { 
		return nil, errors.New("Missing required fields")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password:", err.Error())
		return nil, err
	}

	forgotPasswordUserData := &model.ChangePasswordUserData{
		NewPassword: string(hashedPassword),
	}

	db := config.DB
	userRepo := repository.NewUserRepo(db)

	if err := userRepo.ChangePassword(ctx, forgotPasswordUserData, req.OldPassword, req.Id); err != nil {
		log.Println("Failed to reset password:", err.Error())
		return nil, err
	}

	return &pb.ChangePasswordResponse{Message: "Password resetted!"}, nil
}

func (s *UserServiceServer) UpdateDistanceTravelled(ctx context.Context, req *pb.UpdateDistanceTravelledRequest) (*pb.UpdateDistanceTravelledResponse, error) {
	// Validating the request inputs
	if req.Id == 0 {
		return nil, errors.New("Id is required")
	}

	if req.Distance <= 0 {
		return nil, errors.New("Distance must be a positive number")
	}

	updateDistanceUserData := &model.UpdateDistanceUserData{
		Distance: req.Distance,
	}

	db := config.DB
	userRepo := repository.NewUserRepo(db)

	// Updating the user's distance travelled in the database
	if err := userRepo.UpdateDistanceTravelled(ctx, updateDistanceUserData, uint64(req.Id)); err != nil {
		log.Println("Failed to update distance travelled:", err.Error())
		return nil, err
	}

	return &pb.UpdateDistanceTravelledResponse{Message: "Distance travelled updated!"}, nil
}

func (s *UserServiceServer) AuthenticateUser(ctx context.Context, req *pb.AuthenticateUserRequest) (*pb.AuthenticateUserResponse, error) {
	if req.Token == "" {
		return &pb.AuthenticateUserResponse{Valid: false, Message: "Token is required"}, errors.New("Token is required")
	}

	parsedId, err := parseToken(req.Token, os.Getenv("JWT_SECRET"))
	if err != nil {
		log.Println("Failed to parse token:", err.Error())
		return &pb.AuthenticateUserResponse{Valid: false, Message: err.Error()}, err
	}
	log.Printf("Extracted claims ID: %v", parsedId)

	user := &model.User{}
	err = config.DB.Model(&model.User{}).Where("id = ?", parsedId).First(user).Error
	if err != nil || reflect.DeepEqual(user, &pb.User{}) {
		log.Println("Failed to get users:", err.Error())
		return &pb.AuthenticateUserResponse{Valid: false, Message: "Invalid Credentials!"}, errors.New("Invalid credentials")
	}

	log.Printf("User found with ID: %v", user.Id)
	return &pb.AuthenticateUserResponse{Valid: true, Message: "Authenticated!", UserId: uint64(user.Id)}, nil
}

func generateToken(user *model.User) (string, error) {
	// Creating the token with user ID and expiration time
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 2).Unix(),
	})
	
	// Sign the token and return it as a string
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func parseToken(tokenString, secret string) (int64, error) {
	// Parsing the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate that the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		log.Println("Error parsing token:", err)
		return 0, err
	}

	// Extract claims and validate them
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// Verify token expiration
		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return 0, errors.New("token expired")
		}

		// Convert the "id" claim to int64
		idFloat, ok := claims["id"].(float64)
		if !ok {
			return 0, errors.New("invalid ID format in token")
		}

		return int64(idFloat), nil
	}
	return 0, errors.New("Invalid token")
}

// func GetUserById(db *gorm.DB) func(c *gin.Context) {
// 	return func(c *gin.Context) {
// 		var data model.User

// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})

// 			return
// 		}
		
// 		data.Id = id
// 		// Connect to DB
// 		userRepo := repository.NewUserRepo(db)
// 		if err := userRepo.GetUserById(c.Request.Context(), &data); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})

// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{
// 			"data": data,
// 		})
// 	}
// }

// func UpdateUser(db *gorm.DB) func(c *gin.Context) {
// 	return func(c *gin.Context) {
// 		var data model.UserUpdate

// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})

// 			return
// 		}

// 		if err := c.ShouldBind(&data); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})

// 			return
// 		}
		
// 		// Connect to DB
// 		userRepo := repository.NewUserRepo(db)
// 		if err := userRepo.UpdateUser(c.Request.Context(), &data, id); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})

// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{
// 			"data": true,
// 		})
// 	}
// }
