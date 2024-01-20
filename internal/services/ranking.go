package services

import (
	"github.com/bwmarrin/discordgo"
	"github.com/leoff00/ta-pago-bot/pkg/helpers"
)

func ExecuteRankingService(i *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	discordData := helpers.GetUserData(i)

	msgEmbed := &discordgo.MessageEmbed{
		Title:       "Ranking dos mais saúdaveis e marombeiros. 💪🏅",
		Description: discordData.Id,
		Type:        discordgo.EmbedTypeRich,
		Color:       10,
	}

	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{msgEmbed},
	}
}
