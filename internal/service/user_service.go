package service

import (
	"context"
	"encoding/json"
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

type AuthClaims struct {
	Id uint64 `json:"id,omitempty"`
	jwt.RegisteredClaims
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

	signUpData := &model.UserData{
		Name:      req.Name,
		PhoneNumber: req.PhoneNumber,
		Email: req.Email,
		Password:  string(hashedPassword),
	}

	db := config.DB
	userRepo := repository.NewUserRepo(db)

	user, err := userRepo.SignUp(ctx, signUpData)

	if err != nil {
		log.Println("Failed to signup:", err.Error())
		return nil, err
	}

	token, err := generateToken(user)
	if err != nil {
		log.Println("Failed to generate token:", err.Error())
		return nil, errors.New("Invalid Phone Number or Password")
	}

	return &pb.SignUpResponse{Token: token}, nil
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

func (s *UserServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	if req.Id == 0 || req.Name == "" || req.PhoneNumber == "" || req.Email == "" || req.Password == "" {
		return nil, errors.New("Missing required fields")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password:", err.Error())
		return nil, err
	}

	updateData := &model.UserData{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Email: req.Email,
		Password: string(hashedPassword),
	}

	db := config.DB
	userRepo := repository.NewUserRepo(db)

	if err := userRepo.UpdateUser(ctx, updateData, int64(req.Id)); err != nil {
		log.Println("Failed to update user:", err.Error())
		return nil, err
	}

	return &pb.UpdateUserResponse{Message: "User updated!"}, nil
}

// func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
// 	var user model.User;
// 	if req.Id == 0 {
// 		return nil, errors.New("Id is required")
// 	}

// 	user.Id = req.Id

// 	db := config.DB
// 	userRepo := repository.NewUserRepo(db)

// 	user, err := userRepo.GetUser(ctx, &user)

// 	if err != nil {
// 		log.Println("Failed to get user:", err.Error())
// 		return nil, err
// 	}

// 	// Marshalling...
// }

func (s *UserServiceServer) AuthenticateUser(ctx context.Context, req *pb.AuthenticateUserRequest) (*pb.AuthenticateUserResponse, error) {
	if req.Token == "" {
		return &pb.AuthenticateUserResponse{Valid: false, Message: "Token is required"}, errors.New("Token is required")
	}

	claims, err := parseToken(req.Token, os.Getenv("JWT_SECRET"))
	if err != nil {
		log.Println("Failed to parse token:", err.Error())
		return &pb.AuthenticateUserResponse{Valid: false, Message: err.Error()}, err
	}

	user := &model.User{}
	err = config.DB.Model(&model.User{}).Where("id = ?", claims.Id).First(user).Error
	if err != nil || reflect.DeepEqual(user, &pb.User{}) {
		log.Println("Failed to get users:", err.Error())
		return &pb.AuthenticateUserResponse{Valid: false, Message: "Invalid Credentials!"}, errors.New("Invalid credentials")
	}

	return &pb.AuthenticateUserResponse{Valid: true, Message: "Authenticated!", UserId: uint64(user.Id)}, nil
}

func generateToken(user *model.User) (string, error) {
	now := time.Now()
	expiry := time.Now().Add(time.Hour * 24 * 2)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthClaims{
		Id: uint64(user.Id),
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"eco-taxi-user-service"},
			ExpiresAt: jwt.NewNumericDate(expiry),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func parseToken(tokenString, secret string) (claims AuthClaims, err error) {
	decodedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return AuthClaims{}, err
	}

	if claims, ok := decodedToken.Claims.(jwt.MapClaims); ok && decodedToken.Valid &&
		claims.VerifyAudience("eco-taxi-user-service", true) &&
		claims.VerifyExpiresAt(time.Now().Unix(), true) &&
		claims.VerifyIssuedAt(time.Now().Unix(), true) {

		authClaims := AuthClaims{}
		b, err := json.Marshal(claims)
		if err != nil {
			return AuthClaims{}, err
		}
		err = json.Unmarshal(b, &authClaims)
		if err != nil {
			return AuthClaims{}, err
		}
		return authClaims, nil
	}
	return AuthClaims{}, err
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
