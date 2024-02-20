package main

import (
	"fmt"
	"github.com/leoff00/ta-pago-bot/internal/bot"
	"github.com/leoff00/ta-pago-bot/internal/repo"
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
	db := setup.DB()
	
	repository := repo.NewUserRepository(db)
	service := services.NewActivitiesServices(repository)
	cron := services.NewCronService(repository, service)

	fmt.Println(displayArt)
	bot.Start(service, cron)
}
