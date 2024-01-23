package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/leoff00/ta-pago-bot/internal/services"
	"github.com/leoff00/ta-pago-bot/pkg/helpers"
)

var (
	activities = services.ActivitiesServices{}
)

func (ih *InteractionsHandlers) Join() InteractionCreateResponse {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == AppCmd {
			switch i.ApplicationCommandData().Name {
			case "inscrever":
				joinResponse := activities.ExecuteJoinService(i)
				helpers.InteractionResponseFactory(joinResponse, s, i)
			}
		}
	}
}

func (ih *InteractionsHandlers) Pay() InteractionCreateResponse {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == AppCmd {
			switch i.ApplicationCommandData().Name {
			case "ta-pago":
				payResponse := activities.ExecutePayService(i)
				helpers.InteractionResponseFactory(payResponse, s, i)
			}
		}
	}
}

func (ih *InteractionsHandlers) Ranking() InteractionCreateResponse {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == AppCmd {
			switch i.ApplicationCommandData().Name {
			case "ranking":
				rankingResponse, _ := activities.ExecuteRankingService()
				helpers.InteractionResponseFactory(rankingResponse, s, i)
			}
		}
	}
}

func (ih *InteractionsHandlers) Help() InteractionCreateResponse {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == AppCmd {
			switch i.ApplicationCommandData().Name {
			case "help":
				helpResponse := activities.HelpCmd()
				helpers.InteractionResponseFactory(helpResponse, s, i)
			}
		}
	}
}

func ExecHandlers(bot *discordgo.Session) {
	ih := InteractionsHandlers{}
	bot.AddHandlerOnce(OnReady())
	bot.AddHandler(ih.Join())
	bot.AddHandler(ih.Pay())
	bot.AddHandler(ih.Ranking())
	bot.AddHandler(ih.Help())
}
