package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	
)


type Config struct {
	AppPort string
	AppEnv string
}


func LoadConfig() *Config {      
    err := godotenv.Load()
    if err != nil {
        log.Fatal("error loading .env file")
    }


	return &Config{
		AppPort:getenv("APP_PORT","8080"),
		AppEnv: getenv("APP_ENV","development"),
	}
}


func getenv(key , defaultValue string) string {
	if value , exists := os.LookupEnv(key) ; exists {
		return value
	}
	return defaultValue
}