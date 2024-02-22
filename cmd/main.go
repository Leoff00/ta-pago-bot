package main

import (
	"fmt"
	"github.com/leoff00/ta-pago-bot/internal/bot"
	"github.com/leoff00/ta-pago-bot/internal/services"
	"github.com/leoff00/ta-pago-bot/internal/setup"
	"github.com/leoff00/ta-pago-bot/pkg/env"
	"github.com/leoff00/ta-pago-bot/pkg/timezone"
)

const displayArt = "\033[36m" + ` 
$$$$$$$$\  $$$$$$\        $$$$$$$\   $$$$$$\   $$$$$$\   $$$$$$\  $$\ 
\__$$  __|$$  __$$\       $$  __$$\ $$  __$$\ $$  __$$\ $$  __$$\ $$ |
   $$ |   $$ /  $$ |      $$ |  $$ |$$ /  $$ |$$ /  \__|$$ /  $$ |$$ |
   $$ |   $$$$$$$$ |      $$$$$$$  |$$$$$$$$ |$$ |$$$$\ $$ |  $$ |$$ |
   $$ |   $$  __$$ |      $$  ____/ $$  __$$ |$$ |\_$$ |$$ |  $$ |\__|
   $$ |   $$ |  $$ |      $$ |      $$ |  $$ |$$ |  $$ |$$ |  $$ |    
   $$ |   $$ |  $$ |      $$ |      $$ |  $$ |\$$$$$$  | $$$$$$  |$$\ 
   \__|   \__|  \__|      \__|      \__|  \__| \______/  \______/ \__|
                                 The bot.
                                  ❚█══█❚
` + "\033[0m"

func main() {
	defaultsEnvs := map[string]string{
		"TZ_APP": "America/Sao_Paulo",
		"ENV":    "DEV",
	}
	env.Load(defaultsEnvs, ".env")
	setup.CheckEnvs()
	environment := env.Getenv("ENV")
	timezone.Load(env.Getenv("TZ_APP"))

	tenantCfg := setup.Tenants()
	service := setup.Service(tenantCfg, environment)
	cron := services.NewCronService(service)

	fmt.Println(displayArt)
	bot.Start(service, cron)
}
