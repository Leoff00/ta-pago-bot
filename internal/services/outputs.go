package services

import (
	"github.com/bwmarrin/discordgo"
)

/*
failOutput creates an EXPECTED business service fail output.

	Notes:
		To unexpected errors, use failUnexpectedOutput instead.

	Defaults:
		'Title: Deu merda aqui!!!'
*/
func failOutput(opts OutOpt) *discordgo.InteractionResponseData {
	title := "Deu merda aqui!!!"
	if opts.Title != "" {
		title = opts.Title
	}
	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{
			&discordgo.MessageEmbed{
				Title:       title,
				Description: opts.Description,
				Type:        discordgo.EmbedTypeRich,
				Color:       10,
			}},
	}
}

// successOutput creates expected service output.
func successOutput(opts OutOpt) *discordgo.InteractionResponseData {
	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{
			&discordgo.MessageEmbed{
				Title:       opts.Title,
				Description: opts.Description,
				Type:        discordgo.EmbedTypeRich,
				Color:       10,
			}},
	}
}

// failUnexpectedOutput creates an UNEXPECTED business service fail output.
func failUnexpectedOutput() *discordgo.InteractionResponseData {
	return failOutput(OutOpt{
		Title:       "Deu merda aqui!!!",
		Description: "Ocorreu um erro inesperado. Verifica essa parada ai mermão!",
	})
}

/*
OutOpt is a option struct for successOutput & failOutput options

	Usages:
		successOutput(OutOpt{
			Title:       "Agora é só mandar bala",
			Description: "Digite o comando /ta-pago toda vez que buscar o shape meu nobre!! 💪🏅",
		})
		failOutput(OutOpt{
			Description: "Não foi possível criar o usuário",
		})
*/
type OutOpt struct {
	Title       string
	Description string
}
