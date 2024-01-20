package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/leoff00/ta-pago-bot/internal/services"
	"github.com/leoff00/ta-pago-bot/pkg/helpers"
)

func join() InteractionCreateResponse {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		joinResponse := services.ExecuteJoinService(i)
		if i.Type == AppCmd {
			switch i.ApplicationCommandData().Name {
			case "inscrever":
				helpers.InteractionResponseFactory(joinResponse, s, i)
			}
		}
	}
}

func pay() InteractionCreateResponse {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		payResponse := services.ExecutePayService(i)
		if i.Type == AppCmd {
			switch i.ApplicationCommandData().Name {
			case "ta-pago":
				helpers.InteractionResponseFactory(payResponse, s, i)
			}
		}
	}
}

// func ranking() InteractionCreateResponse {
// 	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 		rankingResponse := services.ExecuteRankingService(i)
// 		if i.Type == AppCmd {
// 			switch i.ApplicationCommandData().Name {
// 			case "ranking":
// 				helpers.InteractionResponseFactory(rankingResponse, s, i)
// 			}
// 		}
// 	}
// }

// func help() InteractionCreateResponse {
// 	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 		helpResponse := services.HelpCmd(i)

// 		if i.Type == AppCmd {
// 			switch i.ApplicationCommandData().Name {
// 			case "help":
// 				helpers.InteractionResponseFactory(helpResponse, s, i)
// 			}
// 		}
// 	}
// }

func ExecHandlers(bot *discordgo.Session) {
	bot.AddHandlerOnce(OnReady())
	bot.AddHandler(join())
	bot.AddHandler(pay())
	// bot.AddHandler(ranking())
	// bot.AddHandler(help())
}
