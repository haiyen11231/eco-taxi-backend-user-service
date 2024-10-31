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
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (userRepo *userRepo) SignUp(ctx context.Context, data *model.SignUpUserData) error {
	if err := userRepo.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (userRepo *userRepo) LogIn(ctx context.Context, data *model.LogInUserData) (*model.User, error) {
	var user model.User
	
	if err := userRepo.db.Where("phone_number = ?", data.PhoneNumber).First(&user).Error; err != nil {
		log.Println("Failed to get user by username:", err.Error())
		return nil, errors.New("Invalid Phone Number or Password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		log.Println("Failed to compare password:", err.Error())
		return nil, errors.New("Invalid Phone Number or Password")
	}

	return &user, nil
}

func (userRepo *userRepo) ForgotPassword(ctx context.Context, data *model.ChangePasswordUserData, email string) error {
	if err := userRepo.db.Where("email = ?", email).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (userRepo *userRepo) UpdateUser(ctx context.Context, data *model.UpdateUserData, id uint64) error {
	if err := userRepo.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (userRepo *userRepo) GetUser(ctx context.Context, data *model.User) error {
	if err := userRepo.db.First(&data).Error; err != nil {
		return err
	}

	return nil
}

func (userRepo *userRepo) ChangePassword(ctx context.Context, data *model.ChangePasswordUserData, oldPassword string, id uint64) error {
	var user model.User
	
	if err := userRepo.db.Where("id = ?", id).First(&user).Error; err != nil {
		log.Println("Failed to get user by id:", err.Error())
		return errors.New("Invalid Id")
	}

	// Check oldPassword
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		log.Println("Failed to compare password:", err.Error())
		return errors.New("Invalid Password")
	}

	// Change password
	if err := userRepo.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (userRepo *userRepo) UpdateDistanceTravelled(ctx context.Context, data *model.UpdateDistanceUserData, id uint64) error {
	var user model.User;

	// Retrieving user to confirm existence
	if err := userRepo.db.Where("id = ?", id).First(&user).Error; err != nil {
		log.Println("Failed to get user by id:", err.Error())
		return errors.New("Invalid Id")
	}

	data.Distance += user.DistanceTravelled

	if err := userRepo.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}