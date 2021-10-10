package main

import (
	"cager/Config"
	"cager/Models"
	"cager/Routes"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(postgres.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{})

	if err != nil {
		fmt.Println("Status DB:", err)
	}

	db.AutoMigrate(
		&Models.User{},
		&Models.Category{},
		&Models.Product{},
		&Models.Page{},
		&Models.Slider{},
		&Models.SliderRelation{},
		&Models.CategoryRelation{},
		&Models.Setting{})

	router := Routes.SetupRouter(db)
	router.Run(":" + os.Getenv("APP_PORT"))

}
