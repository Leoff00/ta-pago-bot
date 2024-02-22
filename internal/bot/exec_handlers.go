package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/leoff00/ta-pago-bot/internal/services"
	"github.com/leoff00/ta-pago-bot/pkg/discord"
)

type InteractionsHandlers struct {
	services *services.ActivitiesServices
}

func (ih *InteractionsHandlers) join() InteractionCreateResponse {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == AppCmd {
			switch i.ApplicationCommandData().Name {
			case "inscrever":
				joinResponse := ih.services.ExecuteJoin(i)
				discord.InteractionResponseFactory(joinResponse, s, i)
			}
		}
	}
}

func (ih *InteractionsHandlers) pay() InteractionCreateResponse {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == AppCmd {
			switch i.ApplicationCommandData().Name {
			case "ta-pago":
				payResponse := ih.services.ExecutePay(i)
				discord.InteractionResponseFactory(payResponse, s, i)
			}
		}
	}
}

func (ih *InteractionsHandlers) ranking() InteractionCreateResponse {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == AppCmd {
			switch i.ApplicationCommandData().Name {
			case "ranking":
				rankingResponse := ih.services.ExecuteRanking(i, "")
				discord.InteractionResponseFactory(rankingResponse, s, i)
			}
		}
	}
}

func (ih *InteractionsHandlers) reset() InteractionCreateResponse {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == AppCmd {
			switch i.ApplicationCommandData().Name {
			case "reset":
				restartResponse := ih.services.ExecuteReset(i)
				discord.InteractionResponseFactory(restartResponse, s, i)
			}
		}
	}
}

func (ih *InteractionsHandlers) help() InteractionCreateResponse {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == AppCmd {
			switch i.ApplicationCommandData().Name {
			case "help":
				helpResponse := ih.services.HelpCmd()
				discord.InteractionResponseFactory(helpResponse, s, i)
			}
		}
	}
}
