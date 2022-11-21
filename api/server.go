package api

import (
	"log"
	"os"

	"api-court-rental/api/controllers"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error getting env: %v", err)
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	server.Run("127.0.0.1:" + os.Getenv("PORT"))
}
