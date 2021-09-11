package main

import (
	"cager/auth"
	"cager/category"
	"cager/handler"
	"cager/middleware"
	"cager/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=andre dbname=exhotel port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// REPOSITORY
	userRepository := user.NewRepository(db)
	categoryRepository := category.NewRepository(db)

	// SERVICE
	userService := user.NewService(userRepository)
	categoryService := category.NewService(categoryRepository)

	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	categoryHandler := handler.NewCategoryHandler(categoryService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/user/email/checker", userHandler.ChekEmailAvailability)
	// AFTER LOGIN
	api.POST("/user/change/name", middleware.AuthMiddleware(authService, userService), userHandler.ChangeNameHandler)
	api.GET("/user/detail", middleware.AuthMiddleware(authService, userService), userHandler.FetchUser)
	api.GET("/users", middleware.AuthMiddleware(authService, userService), userHandler.GetAllUsers)
	api.POST("/user/change/password", middleware.AuthMiddleware(authService, userService), userHandler.ChangePassword)
	api.POST("/user/delete", middleware.AuthMiddleware(authService, userService), userHandler.DeleteUser)

	// CATEGORY
	api.GET("/category", categoryHandler.GetAllCategory)
	api.GET("/category/:id", categoryHandler.GetCategoryById)
	api.GET("/category/by/:slug", categoryHandler.GetCategoryBySlug)
	api.POST("/category/register", middleware.AuthMiddleware(authService, userService), categoryHandler.RegisterCategory)
	api.POST("/category/update", middleware.AuthMiddleware(authService, userService), categoryHandler.UpdateCategory)
	api.POST("/category/delete", middleware.AuthMiddleware(authService, userService), categoryHandler.DeleteCategory)

	router.Run(":8090")

	// http.ListenAndServe(":8090", nil)

}
