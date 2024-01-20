package services

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/leoff00/ta-pago-bot/internal/repo"
	"github.com/leoff00/ta-pago-bot/pkg/helpers"
)

func ExecuteJoinService(i *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	du := helpers.GetUserData(i)

	if err := repo.Save(du); err != nil {
		msgEmbed := &discordgo.MessageEmbed{
			Title:       err.Error(),
			Description: "Opa! Parece que você tentou se inscrever mais de uma vez...",
			Type:        discordgo.EmbedTypeRich,
			Color:       10,
		}

		return &discordgo.InteractionResponseData{
			Embeds: MsgEmbedType{msgEmbed},
		}
	}

	titleFmt := fmt.Sprintf("%s acabou de se inscrever na lista do TA PAGO!", du.Username)
	msgEmbed := &discordgo.MessageEmbed{
		Title:       titleFmt,
		Description: "Para contabilizar na lista, digite o comando /ta-pago toda vez que pagar um treino! 💪🏅",
		Type:        discordgo.EmbedTypeRich,
		Color:       10,
	}

	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{msgEmbed},
	}
}
