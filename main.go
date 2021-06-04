package main

import (
	"cager/auth"
	"cager/balance"
	"cager/handler"
	"cager/middleware"
	"cager/payment"
	"cager/topup"
	"cager/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=andre dbname=ganeshdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	paymentRepository := payment.NewRepository(db)
	balanceRepository := balance.NewRepository(db)
	topupRepository := topup.NewRepository(db)

	userService := user.NewService(userRepository)
	paymentService := payment.NewService(paymentRepository)
	balanceService := balance.NewService(balanceRepository, userService, paymentService, userRepository)
	topupService := topup.NewService(topupRepository)

	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	paymentHandler := handler.NewPaymentHandler(paymentService)
	balanceHandler := handler.NewBalanceHandler(balanceService, topupService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email-checkers", userHandler.ChekEmailAvailability)
	api.POST("/phone-checkers", userHandler.ChekPhoneAvailability)

	// PAYMENT
	api.POST("/payment-register", middleware.AuthMiddleware(authService, userService), paymentHandler.RegisterPayment)
	api.GET("/payments", paymentHandler.Index)

	// BALANCE
	api.POST("/balance-topup", middleware.AuthMiddleware(authService, userService), balanceHandler.CreateBalance)
	api.POST("/topup-approve", middleware.DBApproveBalanceMiddleware(db), balanceHandler.BalanceApprove)

	// PROFILE
	api.POST("/avatars", middleware.AuthMiddleware(authService, userService), userHandler.UploadAvatar)
	api.POST("/change-name", middleware.AuthMiddleware(authService, userService), userHandler.ChangeName)

	api.POST("/check-pin", middleware.AuthMiddleware(authService, userService), userHandler.HandlerCheckPin)
	api.POST("/check-pin-temporary", middleware.AuthMiddleware(authService, userService), userHandler.HandlerCheckPinTemporary)
	api.POST("/change-pin", middleware.AuthMiddleware(authService, userService), userHandler.HandlerChangePin)
	api.POST("/change-pin-temporary", middleware.AuthMiddleware(authService, userService), userHandler.HandlerChangePinTemporary)

	api.POST("/change-phone-number", middleware.AuthMiddleware(authService, userService), userHandler.HandlerChangePhoneNumber)
	api.POST("/change-email", middleware.AuthMiddleware(authService, userService), userHandler.ChangeEmailHandler)
	router.Run()

}
