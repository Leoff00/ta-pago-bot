package env

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	defaultsSet bool
)

func Getenv(env string) string {
	if !defaultsSet {
		LoadDefaults()
	}
	if val := viper.GetString(env); val != "" {
		return val
	}
	osEnv := os.Getenv(env)
	if osEnv == "" {
		log.Default().Fatalf("Environment variable %s not set and no default available", env)
	}
	return osEnv
}
func LoadDefaults() {
	viper.SetDefault("DB_PATH", "./db")
	viper.SetDefault("DB_NAME", "ta_pago.db")
	viper.SetDefault("TZ_BOT", "-03") // default as "America/Sao_Paulo"
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Default().Printf("Not .env file found : %v\n", err)
	}
	defaultsSet = true
	logEnvs()
}

func logEnvs() {
	defaults := viper.AllSettings()
	concatenateDefaultsToString := ""
	for k, v := range defaults {
		k = strings.ToUpper(k)
		if k == "TOKEN" {
			concatenateDefaultsToString += fmt.Sprintf("%s: %s | ", k, "**********")
			continue
		}
		concatenateDefaultsToString += fmt.Sprintf("%s: %s | ", k, v)
	}
	log.Default().Println(concatenateDefaultsToString)
}
