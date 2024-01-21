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
	fmt.Println(rank)
	var top1, top2, top3, top4, top5, top6, top7, top8, top9, top10 string
	var c1, c2, c3, c4, c5, c6, c7, c8, c9, c10 int
	top1, c1 = rank[0].Username, rank[0].Count
	top2, c2 = rank[1].Username, rank[1].Count
	top3, c3 = rank[2].Username, rank[2].Count
	top4, c4 = rank[3].Username, rank[3].Count
	top5, c5 = rank[4].Username, rank[4].Count
	top6, c6 = rank[5].Username, rank[5].Count
	top7, c7 = rank[6].Username, rank[6].Count
	top8, c8 = rank[7].Username, rank[7].Count
	top9, c9 = rank[8].Username, rank[8].Count
	top10, c10 = rank[9].Username, rank[9].Count

	fmtDescription := fmt.Sprintf(`
	## LISTA DOS 10 PRIMEIROS MAIS EXERCITADOS DO SERVER: 
	TOP 1 %-15s -  	 %d 🥇🏆
	TOP 2 %-15s -    %d 🥈🏆
	TOP 3 %-15s -    %d 🥉🏆
	TOP 4 %-15s -  	 %d
	TOP 5 %-15s -  	 %d
	TOP 6 %-15s -  	 %d
	TOP 7 %-15s -  	 %d
	TOP 8 %-15s -    %d
	TOP 9 %-15s -    %d
	TOP 10 %-15s -    %d`,
		top1, c1, top2, c2, top3, c3, top4, c4, top5, c5, top6, c6, top7, c7, top8, c8, top9, c9, top10, c10)
	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{
			&discordgo.MessageEmbed{
				Title:       "Ranking dos mais saúdaveis e marombeiros. 💪🏅",
				Description: fmtDescription,
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
