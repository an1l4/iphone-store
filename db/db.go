package db

import (
	"log"
	"os"

	"github.com/an1l4/iphoneshop/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB{
	err:=godotenv.Load(".env")

	if err!=nil{
		log.Fatal("error in loading .env file")
	}

	db,err:=gorm.Open(postgres.Open(os.Getenv("POSTGRES_URL")),&gorm.Config{})

	if err!=nil{
		panic(err)
	}

	db.AutoMigrate(&models.Iphone{})

	return db

}