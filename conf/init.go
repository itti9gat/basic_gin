package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ServerPort    = ""
	MySqlHost     = ""
	MySqlPort     = ""
	MySqlUser     = ""
	MySqlPassword = ""
	MySqlName     = ""
)

func init() {

	currentEnvironment, _ := os.LookupEnv("ENVIRONMENT")

	if currentEnvironment != "test" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	ServerPort = os.Getenv("SERVER_PORT")
	MySqlHost = os.Getenv("DB_HOST")
	MySqlPort = os.Getenv("DB_PORT")
	MySqlUser = os.Getenv("DB_USER")
	MySqlPassword = os.Getenv("DB_PASSWORD")
	MySqlName = os.Getenv("DB_NAME")
}
