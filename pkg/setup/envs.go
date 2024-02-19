package setup

import (
	"github.com/leoff00/ta-pago-bot/pkg/env"
	"log"
)

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
