package main

import (
	"fmt"
	"github.com/leoff00/ta-pago-bot/internal/bot"
	"github.com/leoff00/ta-pago-bot/internal/services"
	"github.com/leoff00/ta-pago-bot/pkg/env"
	"github.com/leoff00/ta-pago-bot/pkg/setup"
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
	setup.Envs()
	setup.Pwd()
	setup.TimeZone(env.Getenv("TZ_BOT"))
	tenantCfg := setup.Tenants()
	service := setup.Service(tenantCfg)
	cron := services.NewCronService(service)

	fmt.Println(displayArt)
	bot.Start(service, cron)
}
