package repository

import (
	"context"
	"errors"
	"log"

	// "github.com/go-redis/redis"
	"github.com/haiyen11231/eco-taxi-backend-user-service/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
	// rdb *redis.Client
}

// func NewUserRepo(db *gorm.DB, rdb *redis.Client) *userRepo {
// 	return &userRepo{
// 		db: db,
// 		rdb: rdb,
// 	}
// }
func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (userRepo *userRepo) SignUp(ctx context.Context, data *model.UserData) error {
	if err := userRepo.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (userRepo *userRepo) LogIn(ctx context.Context, data *model.LogInUserData) (*model.User, error) {
	var user model.User
	
	if err := userRepo.db.Where("phone_number = ?", data.PhoneNumber).First(&user).Error; err != nil {
		log.Println("Failed to get users by username:", err.Error())
		return nil, errors.New("Invalid Phone Number or Password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		log.Println("Failed to compare password:", err.Error())
		return nil, errors.New("Invalid Phone Number or Password")
	}

	return &user, nil
}

func (userRepo *userRepo) UpdateUser(ctx context.Context, data *model.UserData, id int64) error {
	if err := userRepo.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

// err = config.DB.Where("id = ? and user_id = ?", req.Id, req.UserId).Updates(product).Error


func (userRepo *userRepo) GetUser(ctx context.Context, data *model.User) error {
	if err := userRepo.db.First(&data).Error; err != nil {
		return err
	}

	return nil
}