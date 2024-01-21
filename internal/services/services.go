package services

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/leoff00/ta-pago-bot/internal/repo"
	"github.com/leoff00/ta-pago-bot/pkg/helpers"
)

var (
	dur = repo.DiscordUserRepository{}
)

func (as *ActivitiesServices) ExecuteJoinService(i *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	du := helpers.GetUserData(i)
	if err := dur.Save(du); err != nil {
		return &discordgo.InteractionResponseData{
			Embeds: MsgEmbedType{
				&discordgo.MessageEmbed{
					Title:       "Ops...",
					Description: "Parece que você tentou se inscrever mais de uma vez. ❌",
					Type:        discordgo.EmbedTypeRich,
					Color:       10,
				}},
		}
	}

	titleFmt := fmt.Sprintf("%s acabou de se inscrever na lista do TA PAGO!", du.Username)

	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{
			&discordgo.MessageEmbed{
				Title:       titleFmt,
				Description: "Para contabilizar na lista, digite o comando /ta-pago toda vez que pagar um treino! 💪🏅",
				Type:        discordgo.EmbedTypeRich,
				Color:       10,
			}},
	}
}

func (as *ActivitiesServices) ExecutePayService(i *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	du := helpers.GetUserData(i)
	if err := dur.UpdateCount(du.Id); err != nil {
		return &discordgo.InteractionResponseData{
			Embeds: MsgEmbedType{
				&discordgo.MessageEmbed{
					Title:       "Ops...",
					Description: "Você não pode submeter um treino para o contador sem se increver!! ❌",
					Type:        discordgo.EmbedTypeRich,
					Color:       10,
				}},
		}
	}

	fmtTitle := fmt.Sprintf("%s pagou!!!", du.Username)
	fmtDescription := fmt.Sprintf("%s acabou de submeter um treino!!!", du.Username)

	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{
			&discordgo.MessageEmbed{
				Title:       fmtTitle,
				Description: fmtDescription,
				Type:        discordgo.EmbedTypeRich,
				Color:       10,
			}},
	}
}

func (as *ActivitiesServices) ExecuteRankingService(i *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	rank := dur.GetUsers()
	var res string
	emojis := [3]string{"🥇🏆", "🥈🏆", "🥉🏆"}

	for i, v := range rank[:3] {
		res += fmt.Sprintf("\nTOP %d %s - %d %s", i+1, v.Username, v.Count, emojis[i])
	}
	for i, v := range rank[3:10] {
		res += fmt.Sprintf("\nTOP %d %s - %d", i+4, v.Username, v.Count)
	}

	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{
			&discordgo.MessageEmbed{
				Title:       "Ranking dos mais saudáveis e marombeiros. 💪🏅",
				Description: res,
				Type:        discordgo.EmbedTypeArticle,
				Color:       10,
			}},
	}
}

func (as *ActivitiesServices) HelpCmd(i *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	fmtDescription := fmt.Sprintln(`
		/inscrever: Este comando te incluirá na lista de contagem de treinos o autor do comando. ✅ 

		/ta-pago: Este comando alidara a contagem de treino do autor do comando, aumentando sua posição no ranking. 💪
	
		/ranking: Use este comando para visualizar a lista atualizada dos **10 Primeiros** participantes. 🏆🏅
		`)
	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{
			&discordgo.MessageEmbed{
				Title:       "Veja abaixo como os comandos funcionam.",
				Description: fmtDescription,
				Type:        discordgo.EmbedTypeRich,
				Color:       10,
			}},
	}
}
