package env

import (
	"os"

	"github.com/spf13/viper"
)

func Getenv(envFile string) string {
	viper.SetConfigFile("./.env")
	if err := viper.ReadInConfig(); err != nil {
		return os.Getenv(envFile)
	}
	return viper.GetString(envFile)
}
