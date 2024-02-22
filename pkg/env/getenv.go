package env

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

/*
	Getenv returns the environment variable or the default value if not set

Searches .env file first, then OS Envs, then defaults envs
*/
func Getenv(env string) string {
	if val := viper.GetString(env); val != "" {
		return val
	}
	osEnv := os.Getenv(env)
	if osEnv == "" {
		log.Default().Fatalf("Environment variable %s not set and no default available", env)
	}
	return osEnv
}
func Load(defaults map[string]string, envFiles ...string) {
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
	for _, env := range envFiles {
		viper.SetConfigFile(env)
		if err := viper.ReadInConfig(); err != nil {
			log.Default().Printf("ENV WARN: Can't load %s file.", env)
		}
	}
	logEnvs()
}

func logEnvs() {
	defaults := viper.AllSettings()
	concatenateDefaultsToString := ""
	for k, v := range defaults {
		k = strings.ToUpper(k)
		prefixToHide := "SENSITIVE_"
		if strings.HasPrefix(k, prefixToHide) {
			concatenateDefaultsToString += fmt.Sprintf("%s: %s | ", k, "**********")
			continue
		}
		concatenateDefaultsToString += fmt.Sprintf("%s: %s | ", k, v)
	}
	log.Default().Println(concatenateDefaultsToString)
}
