package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	cmds = []*discordgo.ApplicationCommand{
		{
			Name:        "editar_contagem",
			Description: "⚠️ Edit the count of an user (Mod Exclusive only).⚠️",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "user_id",
					Description: "Which the user id",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "count",
					Description: "Which the count",
					Required:    true,
				},
			},
		},
		{
			Name:        "reset",
			Description: "⚠️ reset the db count (Mod exclusive only).⚠️",
		},
		{
			Name:        "help",
			Description: "Figure out what command do! can use this one a lot!",
		},
		{
			Name:        "inscrever",
			Description: "join in the TA PAGO! meeting and count!",
		},
		{
			Name:        "ta-pago",
			Description: "Submit your workout to the bot count!",
		},
		{
			Name:        "ranking",
			Description: "See the best workout doers.",
		},
	}
)

func addCmds(commands []*discordgo.ApplicationCommand, s *discordgo.Session) error {
	for _, cmd := range commands {
		if _, err := s.ApplicationCommandCreate(s.State.Application.ID, "", cmd); err != nil {
			return err
		}
	}
	return nil
}

func OnReady() func(*discordgo.Session, *discordgo.Ready) {
	return func(s *discordgo.Session, _ *discordgo.Ready) {
		DeleteCommands(s.State.Application.ID)
		if err := addCmds(cmds, s); err != nil {
			log.Default().Println("Cannot add commands - on AddCmd Func ->", err.Error())
		}
	}
}

func DeleteCommands(botId string) func(*discordgo.Session, *discordgo.Ready) {
	return func(s *discordgo.Session, _ *discordgo.Ready) {

		commands, err := s.ApplicationCommands(botId, "")
		if err != nil {
			log.Default().Println("Wasn't possible to load... on Command file ->", err.Error())
			return
		}
		for _, command := range commands {
			if err = s.ApplicationCommandDelete(botId, "", command.ID); err != nil {
				log.Default().Println("Cannot remove the commands on Command file ->", err.Error())
			}
		}
	}
}
