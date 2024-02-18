package services

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"

	"github.com/leoff00/ta-pago-bot/internal/repo"
	"github.com/leoff00/ta-pago-bot/pkg/factory"
	"github.com/leoff00/ta-pago-bot/pkg/strings"
)

var dur = repo.DiscordUserRepository{}

func (as *ActivitiesServices) ExecuteJoinService(
	i *discordgo.InteractionCreate,
) *discordgo.InteractionResponseData {
	du := factory.GetUserData(i)
	if err := dur.Save(du); err != nil {
		fmtDescription := fmt.Sprintf(
			"Parece que o canela seca do <@%s> ta tentando me derrubar, TU JA TA INSCRITO SUA MULA!! ",
			du.Id,
		)
		return &discordgo.InteractionResponseData{
			Embeds: MsgEmbedType{
				&discordgo.MessageEmbed{
					Title:       "Deu merda aqui!!",
					Description: fmtDescription,
					Type:        discordgo.EmbedTypeRich,
					Color:       10,
				},
			},
		}
	}

	fmtDescription := strings.RandomizeJoinPhrases(du.Id)

	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{
			&discordgo.MessageEmbed{
				Title:       "Agora é só mandar bala, digite o comando /ta-pago toda vez que buscar o shape meu nobre!! 💪🏅",
				Description: fmtDescription,
				Type:        discordgo.EmbedTypeRich,
				Color:       10,
			},
		},
	}
}

func (as *ActivitiesServices) ExecutePayService(
	i *discordgo.InteractionCreate,
) *discordgo.InteractionResponseData {
	du := factory.GetUserData(i)
	if err := dur.IncrementCount(du.Id); err != nil {
		return &discordgo.InteractionResponseData{
			Embeds: MsgEmbedType{
				&discordgo.MessageEmbed{
					Title:       "Deu merda aqui!",
					Description: fmt.Sprintln(err.Error() + "❌"),
					Type:        discordgo.EmbedTypeRich,
					Color:       10,
				},
			},
		}
	}

	fmtTitle := fmt.Sprintf("<@%s> pagou!!!", du.Id)
	fmtDescription := strings.RandomizePayPhrases(du.Id)

	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{
			&discordgo.MessageEmbed{
				Title:       fmtTitle,
				Description: fmtDescription,
				Type:        discordgo.EmbedTypeRich,
				Color:       10,
			},
		},
	}
}

func (as *ActivitiesServices) ExecuteRankingService() (*discordgo.InteractionResponseData, *discordgo.MessageEmbed) {
	var irdata *discordgo.InteractionResponseData
	var embed *discordgo.MessageEmbed
	var emojiIter string
	var restIter string

	emojis := [3]string{"🥇🏆", "🥈🏆", "🥉🏆"}
	rank := dur.GetUsers()

	if len(rank) == 0 {
		embed = &discordgo.MessageEmbed{
			Title:       "O ranking ainda está vazio... 💭",
			Description: "Os frangos ainda não submeteram treinos para o contador...",
			Type:        discordgo.EmbedTypeArticle,
			Color:       10,
		}
		irdata = &discordgo.InteractionResponseData{
			Embeds: MsgEmbedType{embed},
		}
	}

	if len(rank) > 0 && len(rank) < 3 {
		embed = &discordgo.MessageEmbed{
			Title:       "Opa! Perai...",
			Description: "É necessário pelo menos ter 3 pessoas pra montar um ranking...",
			Type:        discordgo.EmbedTypeArticle,
			Color:       10,
		}
		irdata = &discordgo.InteractionResponseData{
			Embeds: MsgEmbedType{embed},
		}
	}

	if len(rank) == 3 {
		for i, v := range rank[:3] {
			emojiIter += fmt.Sprintf("\nTOP %d %s - %d %s", i+1, v.Username, v.Count, emojis[i])
		}
		embed = &discordgo.MessageEmbed{
			Title:       "Ranking dos mais saudáveis e marombeiros. 💪🏅",
			Description: emojiIter,
			Type:        discordgo.EmbedTypeArticle,
			Color:       10,
		}

		irdata = &discordgo.InteractionResponseData{
			Embeds: MsgEmbedType{embed},
		}
	}

	if len(rank) > 3 {
		for i, v := range rank[:3] {
			emojiIter += fmt.Sprintf("\nTOP %d %s - %d %s", i+1, v.Username, v.Count, emojis[i])
		}

		for i, v := range rank[3:] {
			restIter += fmt.Sprintf("\nTOP %d %s - %d", i+4, v.Username, v.Count)
		}

		finalStr := emojiIter + restIter
		embed = &discordgo.MessageEmbed{
			Title:       "Ranking dos mais saudáveis e marombeiros. 💪🏅",
			Description: finalStr,
			Type:        discordgo.EmbedTypeArticle,
			Color:       10,
		}

		irdata = &discordgo.InteractionResponseData{
			Embeds: MsgEmbedType{embed},
		}
	}
	return irdata, embed
}

func (as *ActivitiesServices) RestartCount() *discordgo.InteractionResponseData {
	if err := dur.RestartCount(); err != nil {
		log.Default().Println("Cannot restart the the count in database On Service", err.Error())
	}

	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{
			&discordgo.MessageEmbed{
				Title:       "Veja abaixo como os comandos funcionam.",
				Description: "fmtDescription",
				Type:        discordgo.EmbedTypeRich,
				Color:       10,
			},
		},
	}
}

func (as *ActivitiesServices) RestartCmd(
	i *discordgo.InteractionCreate,
) *discordgo.InteractionResponseData {
	du := factory.GetUserData(i)
	fmtDescription := fmt.Sprintf(
		"<@%s> usou o comando para resetar as contagens dos frangos!",
		du.Id,
	)
	if err := dur.ModrestartCount(du.Id); err != nil {
		return &discordgo.InteractionResponseData{
			Embeds: MsgEmbedType{
				&discordgo.MessageEmbed{
					Title:       "Deu merda aqui!!!",
					Description: "O comando só pode ser executado por mods mermão!!",
					Type:        discordgo.EmbedTypeRich,
					Color:       10,
				},
			},
		}
	}

	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{
			&discordgo.MessageEmbed{
				Title:       "TABELA RESETADA!!!",
				Description: fmtDescription,
				Type:        discordgo.EmbedTypeRich,
				Color:       10,
			},
		},
	}
}

func (as *ActivitiesServices) HelpCmd() *discordgo.InteractionResponseData {
	fmtDescription := fmt.Sprintln(`
		/inscrever: Este comando te incluirá na lista de contagem de treinos o autor do comando. ✅ 

		/ta-pago: Este comando validara a contagem de treino do autor do comando, aumentando sua posição no ranking. 💪
	
		/ranking: Use este comando para visualizar a lista atualizada dos **10 Primeiros** participantes. 🏆🏅

		/restart: Este comando é utilizado pelos administradores do servidor para resetar a contagem de treinos caso algo dê problema. 🫡💪
		`)

	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{
			&discordgo.MessageEmbed{
				Title:       "Veja abaixo como os comandos funcionam.",
				Description: fmtDescription,
				Type:        discordgo.EmbedTypeRich,
				Color:       10,
			},
		},
	}
}
