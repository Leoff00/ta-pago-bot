package services

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/leoff00/ta-pago-bot/internal/repo"
	"github.com/leoff00/ta-pago-bot/pkg/helpers"
)

func ExecutePayService(i *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	du := helpers.GetUserData(i)

	if err := repo.UpdateCount(du.Id); err != nil {
		msgEmbed := &discordgo.MessageEmbed{
			Title:       err.Error(),
			Description: "Você não pode submeter um treino para o contador sem se increver!! ❌",
			Type:        discordgo.EmbedTypeRich,
			Color:       10,
		}

		return &discordgo.InteractionResponseData{
			Embeds: MsgEmbedType{msgEmbed},
		}
	}

	fmtTitle := fmt.Sprintf("%s pagou!!!", du.Username)
	fmtDescription := fmt.Sprintf("%s acabou de submeter um treino!!!", du.Username)
	msgEmbed := &discordgo.MessageEmbed{
		Title:       fmtTitle,
		Description: fmtDescription,
		Type:        discordgo.EmbedTypeRich,
		Color:       10,
	}

	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{msgEmbed},
	}
}
