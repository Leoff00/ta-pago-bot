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
		log.Default().Fatalln(`
		Cannot initialize the session
		token authentication failed`,
			err.Error())
		return
	}

	user, err := bot.User("@me")

	if err != nil {
		log.Default().Fatalln("Discord bot user not attached", err.Error())
	}

	botId = user.ID

	ExecHandlers(bot, botId)

	bot.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	bot.Identify.Intents = discordgo.PermissionManageMessages

	DeleteCommands(botId)

	if err = bot.Open(); err != nil {
		log.Default().Println(`
		ERROR during open discord websocket on Start Func ->`,
			err.Error())
	}

	defer bot.Close()

	log.Default().Println("Bot is running - on Start Func")
	fmt.Println("Press Ctrl + C to exit.")

	stsignal := make(chan os.Signal, 1)
	signal.Notify(stsignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-stsignal
}
