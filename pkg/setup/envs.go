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
	if tzBot := env.Getenv("TZ_BOT"); tzBot == "" {
		log.Default().Fatalln("Environment variable 'TZ_BOT' not set")
	}
}
