package repository

import (
	"context"

	"github.com/haiyen11231/eco-taxi-backend-user-service/internal/model"
	"github.com/redis/go-redis"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
	rdb *redis.Client
}

func NewUserRepo(db *gorm.DB, rdb *redis.Client) *userRepo {
	return &userRepo{
		db: db,
		rdb: rdb,
	}
}

func (userRepo *userRepo) CreateUser(ctx context.Context, data *model.UserCreation) error {
	if err := userRepo.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (userRepo *userRepo) GetUserById(ctx context.Context, data *model.User) error {
	if err := userRepo.db.First(&data).Error; err != nil {
		return err
	}

	return nil
}

func (userRepo *userRepo) UpdateUser(ctx context.Context, data *model.UserUpdate, id int) error {
	if err := userRepo.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (userRepo *userRepo) DeleteUser(ctx context.Context, id int) error {
	if err := userRepo.db.Table(model.User{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}