package main

import (
	"log"
	"os"
	"time"

	"github.com/iqbal13/hactiv8tugasdua/config"
	"github.com/iqbal13/hactiv8tugasdua/models"
	"github.com/iqbal13/hactiv8tugasdua/routes"
	"github.com/joho/godotenv"
)

func main() {

	os.Setenv("TZ", "Asia/Jakarta")
	time.LoadLocation("Asia/Jakarta")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed Load Environment")
	}
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Items{}, &models.Order{})

	app := routes.SetupRouter()

	// Run server
	app.Run(os.Getenv("APP_PORT"))
}
