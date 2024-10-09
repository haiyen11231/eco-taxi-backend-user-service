package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/haiyen11231/eco-taxi-backend-user-service/internal/model"
	"github.com/haiyen11231/eco-taxi-backend-user-service/internal/repository"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.UserCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		// Connect to DB
		userRepo := repository.NewUserRepo(db)
		if err := userRepo.CreateUser(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data.Id,
		})
	}
}

func GetUserById(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.User

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		
		data.Id = id
		// Connect to DB
		userRepo := repository.NewUserRepo(db)
		if err := userRepo.GetUserById(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

func UpdateUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.UserUpdate

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		
		// Connect to DB
		userRepo := repository.NewUserRepo(db)
		if err := userRepo.UpdateUser(c.Request.Context(), &data, id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

func DeleteUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		
		// Connect to DB
		userRepo := repository.NewUserRepo(db)
		if err := userRepo.DeleteUser(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}