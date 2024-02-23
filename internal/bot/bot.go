package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/leoff00/ta-pago-bot/internal/services"

	"github.com/bwmarrin/discordgo"
	"github.com/leoff00/ta-pago-bot/pkg/env"
)

var (
	botId string
)

func Start(services *services.ActivitiesServices, ct *services.CronTasks) {
	bot, err := discordgo.New(fmt.Sprintf("Bot %s", env.Getenv("SENSITIVE_TOKEN")))
	if err != nil {
		log.Default().Fatalln(`
		Cannot initialize the session
		token authentication failed`,
			err.Error())
		return
	}
	bot.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.PermissionManageMessages
	user, err := bot.User("@me")
	if err != nil {
		log.Default().Fatalln("Discord bot user not attached", err.Error())
	}
	botId = user.ID

	registerHandlers(bot, services)

	ct.ScheduleTasks(bot)
	ct.Cron.Start()

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

	ct.Cron.Stop()
}

func registerHandlers(bot *discordgo.Session, services *services.ActivitiesServices) {
	ih := InteractionsHandlers{
		services: services,
	}
	bot.AddHandlerOnce(OnReady())
	bot.AddHandler(ih.join())
	bot.AddHandler(ih.pay())
	bot.AddHandler(ih.ranking())
	bot.AddHandler(ih.reset())
	bot.AddHandler(ih.edit())
	bot.AddHandler(ih.help())
}
