package bot

import (
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "inscrever",
			Description: "Join in the TA PAGO! meeting and count!",
		},
	}
)
