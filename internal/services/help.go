package services

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func HelpCmd(i *discordgo.InteractionCreate) *discordgo.InteractionResponseData {

	fmtDescription := fmt.Sprintln(`
		/inscrever: Este comando te incluirá na lista de contagem de treinos o autor do comando. ✅ 

		/ta-pago: Este comando alidara a contagem de treino do autor do comando, aumentando sua posição no ranking. 💪
	
		/ranking: Use este comando para visualizar a lista atualizada dos **10 Primeiros** participantes. 🏆🏅
		`)

	msgEmbed := &discordgo.MessageEmbed{
		Title:       "Veja abaixo como os comandos funcionam.",
		Description: fmtDescription,
		Type:        discordgo.EmbedTypeRich,
		Color:       10,
	}

	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{msgEmbed},
	}
}
