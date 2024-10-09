package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haiyen11231/eco-taxi-backend-user-service/internal/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:user-service@tcp(127.0.0.1:3306)/user_db?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
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

	r := gin.Default()

	v1 := r.Group("/v1")
	users := v1.Group("/users")

	users.POST("", service.CreateUser(db))
	users.GET("/:id", service.GetUserById(db))
	users.PATCH("/:id", service.UpdateUser(db))
	users.DELETE("/:id", service.DeleteUser(db))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}