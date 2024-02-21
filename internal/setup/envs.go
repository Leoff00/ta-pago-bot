package setup

import (
	"github.com/leoff00/ta-pago-bot/pkg/env"
	"log"
)

// CheckEnvs necessary and not-default available envs for the application, also business validation envs
func CheckEnvs() {
	if token := env.Getenv("SENSITIVE_TOKEN"); token == "" {
		log.Default().Fatalln("Environment variable 'TOKEN' not set")
	}
}
