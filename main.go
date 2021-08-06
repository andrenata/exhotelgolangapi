package main

import (
	"cager/auth"
	"cager/balance"
	"cager/handler"
	"cager/middleware"
	"cager/payment"
	"cager/pulsa"
	"cager/topup"
	"cager/transfer"
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

	// TEST
	// postBody, _ := json.Marshal(map[string]string{
	// 	"cmd":      "prepaid",
	// 	"username": "wenolooeK13W",
	// 	"sign":     "b9baafeded7a9fc27f3f78f79fd8623b",
	// })

	// responseBody := bytes.NewBuffer(postBody)

	// //Leverage Go's HTTP Post function to make request
	// resp, err := http.Post("https://api.digiflazz.com/v1/price-list", "application/json", responseBody)
	// if err != nil {
	// 	fmt.Println("Error")
	// }
	// // defer resp.Body.Close()

	// //Read the response body
	// body, err := ioutil.ReadAll(resp.Body)

	// if err != nil {
	// 	fmt.Println("Error")
	// }

	// sb := string(body)
	// fmt.Println(sb)

	// END TEST

	userRepository := user.NewRepository(db)
	paymentRepository := payment.NewRepository(db)
	balanceRepository := balance.NewRepository(db)
	topupRepository := topup.NewRepository(db)
	transferRepository := transfer.NewRepository(db)

	userService := user.NewService(userRepository)
	paymentService := payment.NewService(paymentRepository)
	balanceService := balance.NewService(balanceRepository, userService, paymentService, userRepository)
	topupService := topup.NewService(topupRepository)
	transferService := transfer.NewService(transferRepository)
	pulsaService := pulsa.NewService()

	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	paymentHandler := handler.NewPaymentHandler(paymentService)
	balanceHandler := handler.NewBalanceHandler(balanceService, topupService, transferService)
	pulsaHandler := handler.NewPulsaHandler(pulsaService)

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
	api.POST("/balance-transfer", middleware.DBApproveBalanceMiddleware(db), balanceHandler.BalanceTransfer)

	// PROFILE
	api.POST("/avatars", middleware.AuthMiddleware(authService, userService), userHandler.UploadAvatar)
	api.POST("/change-name", middleware.AuthMiddleware(authService, userService), userHandler.ChangeName)

	api.POST("/check-pin", middleware.AuthMiddleware(authService, userService), userHandler.HandlerCheckPin)
	api.POST("/check-pin-temporary", middleware.AuthMiddleware(authService, userService), userHandler.HandlerCheckPinTemporary)
	api.POST("/change-pin", middleware.AuthMiddleware(authService, userService), userHandler.HandlerChangePin)
	api.POST("/change-pin-temporary", middleware.AuthMiddleware(authService, userService), userHandler.HandlerChangePinTemporary)

	api.POST("/change-phone-number", middleware.AuthMiddleware(authService, userService), userHandler.HandlerChangePhoneNumber)
	api.POST("/change-email", middleware.AuthMiddleware(authService, userService), userHandler.ChangeEmailHandler)

	api.GET("/balance", middleware.AuthMiddleware(authService, userService), userHandler.GetBalanceHandler)
	api.GET("/profile-user", middleware.AuthMiddleware(authService, userService), userHandler.GetUserProfile)

	// PULSA
	api.POST("/pulsa-telkomsel", middleware.AuthMiddleware(authService, userService), pulsaHandler.FindByBrand)
	router.Run(":8090")

	// http.ListenAndServe(":8090", nil)

}
