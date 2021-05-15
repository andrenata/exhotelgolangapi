package main

import (
	"cager/auth"
	"cager/handler"
	"cager/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=andre dbname=koolick port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	//tes token
	// token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.zCGBEiC4n4X5jij4lK4nSEtrbebYxELZ6OfBwdm6CJg")
	// if err != nil {
	// 	fmt.Println("Error")
	// }
	// if token.Valid {
	// 	fmt.Println("VALID")
	// } else {
	// 	fmt.Println("Invalid TOken")
	//}

	//test service
	// input := user.LoginInput{
	// 	Email:    "com.andre@bla.com",
	// 	Password: "password",
	// }
	// user, err := userService.Login(input)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(user.Name)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email-checkers", userHandler.ChekEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)
	router.Run()

	// TEST INPUT
	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Test simpan dari service"
	// userInput.Email = "test@gmail.com"
	// userInput.PhoneNumber = "087860062474"
	// userInput.Password = "Password"

	// userService.RegisterUser(userInput)

	// user := user.User{
	// 	Name:     "test simpan",
	// 	Email:    "andre@blabla.com",
	// 	Password: "passwordtest",
	// }

	// userRepository.Save(user)
}

// func handler(c *gin.Context) {
// 	dsn := "host=localhost user=postgres password=andre dbname=koolick port=5432 sslmode=disable TimeZone=Asia/Shanghai"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	var users []user.User
// 	db.Find(&users)

//c.JSON(200, users)
// 	c.JSON(http.StatusOK, users)

// }

// LAYERING on GIN GOLANG
// input data dari user
// handler mapping input to stract
// service mapping stract to struct user
// responsitory save struct to db
// db
