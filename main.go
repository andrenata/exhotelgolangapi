package main

import (
	"cager/auth"
	"cager/balance"
	"cager/handler"
	"cager/helper"
	"cager/payment"
	"cager/topup"
	"cager/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	api.POST("/phone-checkers", userHandler.ChekPhoneAvailability)

	// PAYMENT
	api.POST("/payment-register", authMiddleware(authService, userService), paymentHandler.RegisterPayment)
	api.GET("/payments", paymentHandler.Index)

	// BALANCE
	// api.POST("/balance-topup", authMiddleware(authService, userService), balanceHandler.CreateBalance)
	// api.POST("/balance-approve", authMiddleware(authService, userService), balanceHandler.BalanceApprove)
	api.POST("/money-transfer", DBTransactionMiddleware(db), balanceHandler.BalanceApprove)

	// PROFILE
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.POST("/change-name", authMiddleware(authService, userService), userHandler.ChangeName)

	api.POST("/check-pin", authMiddleware(authService, userService), userHandler.HandlerCheckPin)
	api.POST("/check-pin-temporary", authMiddleware(authService, userService), userHandler.HandlerCheckPinTemporary)
	api.POST("/change-pin", authMiddleware(authService, userService), userHandler.HandlerChangePin)
	api.POST("/change-pin-temporary", authMiddleware(authService, userService), userHandler.HandlerChangePinTemporary)

	api.POST("/change-phone-number", authMiddleware(authService, userService), userHandler.HandlerChangePhoneNumber)
	api.POST("/change-email", authMiddleware(authService, userService), userHandler.ChangeEmailHandler)
	router.Run()

	// TEST INPUT
	// paymentInput := payment.RegisterPaymentInput{}
	// paymentInput.BankName = "BCA"
	// paymentInput.BankNumber = "1234567"
	// paymentInput.AccountName = "Andre"
	// paymentInput.IsActive = 1

	// userService.RegisterUser(userInput)
	// paymentService.RegisterPayment(paymentInput)

	// payment := payment.Payment{
	// 	BankName : "BCA",
	// 	AccountName : "Andre Nata",
	// 	BankNumber : "1350370591",
	// 	IsActive : 1
	// }
	// paymentRepository.Save(payment)
}

// LAYERING on GIN GOLANG
// input data dari user
// handler mapping input to stract
// service mapping stract to struct user
// responsitory save struct to db
// db

//MIDDLEWARE
//Ambil nilai header Authorization : Bearer tokentoken
//ambil nilai token saja
//validasi token
//kita ambil user_id
//Find user by id di database melalui service
//set context isinya user

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))

		user, err := userService.GetUserbyId(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}

// TRX
//StatusInList -> checks if the given status is in the list
func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

// DBTransactionMiddleware : to setup the database transaction middleware
func DBTransactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		txHandle := db.Begin()
		log.Print("beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Set("db_trx", txHandle)
		c.Next()

		if StatusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			log.Print("committing transactions")
			if err := txHandle.Commit().Error; err != nil {
				log.Print("trx commit error: ", err)
			}
		} else {
			log.Print("rolling back transaction due to status code: ", c.Writer.Status())
			txHandle.Rollback()
		}
	}
}
