package setup

import (
	"github.com/leoff00/ta-pago-bot/pkg/env"
	"log"
)

// Envs checks if the expected environment variables are set
func Envs() {
	if token := env.Getenv("TOKEN"); token == "" {
		log.Default().Fatalln("Environment variable 'TOKEN' not set")
	}
	if dbPath := env.Getenv("DB_PATH"); dbPath == "" {
		log.Default().Fatalln("Environment variable 'DB_PATH' not set")
	}
	if dbName := env.Getenv("DB_NAME"); dbName == "" {
		log.Default().Fatalln("Environment variable 'DB_NAME' not set")
	}
	if tzBot := env.Getenv("TZ_BOT"); tzBot == "" {
		log.Default().Fatalln("Environment variable 'TZ_BOT' not set")
	}
}
