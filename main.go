package main

import (
	"cager/auth"
	"cager/category"
	"cager/handler"
	"cager/middleware"
	"cager/pages"
	"cager/product"
	"cager/settings"
	"cager/user"
	"log"

	"github.com/gin-contrib/cors"
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
	productRepository := product.NewRepository(db)
	pageRepository := pages.NewRepository(db)
	settingRepository := settings.NewRepository(db)

	// SERVICE
	userService := user.NewService(userRepository)
	categoryService := category.NewService(categoryRepository)
	productService := product.NewService(productRepository)
	pageService := pages.NewService(pageRepository)
	settingService := settings.NewService(settingRepository)

	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	categoryHandler := handler.NewCategoryHandler(categoryService, authService)
	productHandler := handler.NewProductHandler(productService)
	pageHandler := handler.NewPageHandler(pageService)
	settingHandler := handler.NewSettingHandler(settingService)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowHeaders = []string{"Authorization", "Content-Type", "X-CSRF-Token"}

	router.Use(cors.New(config))

	// public routes
	router.Static("/storage", "./storage")
	// router.Use(static.Serve("/storage", static.LocalFile("/storage", false)))

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
	api.POST("/user/detail", middleware.AuthMiddleware(authService, userService), userHandler.GetUserByID)
	api.POST("/user/change", middleware.AuthMiddleware(authService, userService), userHandler.ChangeDetailHandler)

	// CATEGORY
	api.GET("/category", categoryHandler.GetAllCategory)
	api.GET("/category/:id", categoryHandler.GetCategoryById)
	api.GET("/category/by/:slug", categoryHandler.GetCategoryBySlug)
	api.POST("/category/register", middleware.AuthMiddleware(authService, userService), categoryHandler.RegisterCategory)
	api.POST("/category/update", middleware.AuthMiddleware(authService, userService), categoryHandler.UpdateCategory)
	api.POST("/category/delete", middleware.AuthMiddleware(authService, userService), categoryHandler.DeleteCategory)

	// PRODUCT
	api.GET("/product", productHandler.GetAllProduct)
	api.POST("/product/create", middleware.AuthMiddleware(authService, userService), productHandler.CreateProductName)
	api.POST("/product/update", middleware.AuthMiddleware(authService, userService), productHandler.UpdateProduct)
	api.POST("/product/del", middleware.AuthMiddleware(authService, userService), productHandler.DelProduct)
	api.POST("/product/detail", middleware.AuthMiddleware(authService, userService), productHandler.FindProductByIDHandler)

	// SLIDER
	api.POST("/slider/all", middleware.AuthMiddleware(authService, userService), productHandler.GetAllSliderHanlder)

	api.POST("/slider/del", middleware.AuthMiddleware(authService, userService), productHandler.DelSlider)

	// DISCOUNT
	api.POST("/discount/create", middleware.AuthMiddleware(authService, userService), productHandler.CreateDiscount)

	// SLIDER REALTION
	api.POST("/sliderrelation/create", middleware.AuthMiddleware(authService, userService), productHandler.CreateSliderRelationHandler)
	api.POST("/sliderrelation/get", middleware.AuthMiddleware(authService, userService), productHandler.GetSliderRelationByProductIDHanlder)
	api.POST("/sliderrelation/del", middleware.AuthMiddleware(authService, userService), productHandler.DelSliderRelationHanlder)

	// CATEGORY RELATION
	api.POST("/categoryrelation/create", middleware.AuthMiddleware(authService, userService), productHandler.CreateCategoryRelationHandler)
	api.POST("/categoryrelation/get", middleware.AuthMiddleware(authService, userService), productHandler.FindCategoryRelationHandler)
	api.POST("/categoryrelation/del", middleware.AuthMiddleware(authService, userService), productHandler.DelCategoryRelationHandler)

	// PAGES
	api.POST("/page/all", middleware.AuthMiddleware(authService, userService), pageHandler.FindAllPage)
	api.POST("/page/create", middleware.AuthMiddleware(authService, userService), pageHandler.CreatePage)
	api.POST("/page/id", middleware.AuthMiddleware(authService, userService), pageHandler.FindByid)
	api.POST("/page/del", middleware.AuthMiddleware(authService, userService), pageHandler.DelPage)
	api.POST("/page/update", middleware.AuthMiddleware(authService, userService), pageHandler.UpdatePage)

	// SETTING
	api.POST("/setting/find", middleware.AuthMiddleware(authService, userService), settingHandler.FindByid)
	api.POST("/setting/update", middleware.AuthMiddleware(authService, userService), settingHandler.UpdateSetting)

	api.POST("/slider/create", middleware.AuthMiddleware(authService, userService), productHandler.CreateSlider)

	router.Run(":8090")

	// http.ListenAndServe(":8090", nil)

}
