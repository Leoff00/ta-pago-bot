package helpers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/leoff00/ta-pago-bot/internal/models"
)

func GetUserData(i *discordgo.InteractionCreate) models.DiscordUser {
	discordUser := models.DiscordUser{
		Id:       i.Member.User.ID,
		Username: i.Member.User.Username,
		Count:    0,
	}
	return discordUser
}
