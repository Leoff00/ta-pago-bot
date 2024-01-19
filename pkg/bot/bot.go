package bot

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/leoff00/ta-pago-bot/pkg/env"
)

var botId string

func Start(db *sql.DB) {
	token := env.Getenv("TOKEN")
	bot, err := discordgo.New(fmt.Sprintf("Bot %s", token))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := bot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	botId = u.ID
	fmt.Println("bot id ->", botId)

	bot.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	bot.Identify.Intents = discordgo.PermissionManageMessages

	err = bot.Open()

	defer bot.Close()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log.Default().Println("Bot is running - on Start Func.")
	fmt.Println("Press Ctrl + C to exit.")

	stsignal := make(chan os.Signal, 1)
	signal.Notify(stsignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-stsignal
}
