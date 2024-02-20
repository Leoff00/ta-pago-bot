package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func InteractionResponseFactory(data *discordgo.InteractionResponseData, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: data,
	}); err != nil {
		log.Default().Println(`
		Error during execution of some handler ->`, err.Error())
		// log the unknown handler
	}
}
