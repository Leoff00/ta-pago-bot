package services

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/leoff00/ta-pago-bot/internal/domain"
	"github.com/leoff00/ta-pago-bot/internal/models"
	"github.com/leoff00/ta-pago-bot/internal/repo"
	"github.com/leoff00/ta-pago-bot/pkg/discord"
	"github.com/leoff00/ta-pago-bot/pkg/env"
	"github.com/leoff00/ta-pago-bot/pkg/helpers"
	"log"
	"slices"
	"strings"
)

type ActivitiesServices struct {
	repository *repo.UserRepository
}

func NewActivitiesServices(repository *repo.UserRepository) *ActivitiesServices {
	return &ActivitiesServices{
		repository: repository,
	}
}

func (as *ActivitiesServices) ExecuteJoin(i *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	discordUser := discord.GetUserData(i)
	isSubscribed, err := as.repository.ExistsById(discordUser.Id)
	if err != nil {
		return failUnexpectedOutput()
	}
	if isSubscribed {
		d := fmt.Sprintf(
			"Parece que o canela seca do %s ta tentando me derrubar, TU JA TA INSCRITO SUA MULA!!",
			discordUser.Member.Mention())
		return failOutput(OutOpt{
			Description: d,
		})
	}
	user, err := domain.NewUser(domain.CreateUserOpts{
		Id:       discordUser.Id,
		Username: discordUser.Username,
		Nickname: discordUser.Nickname,
	})
	if err != nil {
		log.Default().Println("Error during user creation", err.Error())
		return failUnexpectedOutput()
	}
	if err = as.repository.Insert(user); err != nil {
		return failUnexpectedOutput()
	}
	return successOutput(OutOpt{
		Title:       "Agora é só mandar bala, digite o comando /ta-pago toda vez que buscar o shape meu nobre!! 💪🏅",
		Description: helpers.RandomizeJoinPhrases(discordUser.Member.Mention()),
	})
}

func (as *ActivitiesServices) ExecutePay(i *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	discordUsr := discord.GetUserData(i)
	user, err := as.repository.GetUserById(i.Member.User.ID)
	if err != nil {
		return failUnexpectedOutput()
	}
	aggregate := &models.UserAggregate{
		User:        user,
		DiscordUser: discordUsr,
	}
	if user.IsNotSubscribed() {
		return failOutput(OutOpt{
			Description: "você precisa antes se inscrever na lista fera"})
	}
	if user.AlreadySubmitted() {
		return failOutput(OutOpt{
			Description: "seu frango! tu já treinou hoje mermão, volta amanhã"})
	}
	user.Pay()
	err = as.repository.Save(aggregate)
	if err != nil {
		return failUnexpectedOutput()
	}
	return successOutput(OutOpt{
		Title:       fmt.Sprintf("%s pagou!!!", user.GetNickname()),
		Description: helpers.RandomizePayPhrases(discordUsr.Member.Mention()),
	})
}

func (as *ActivitiesServices) ExecuteRanking() *discordgo.InteractionResponseData {
	var emojiIter string
	var restIter string
	emojis := [3]string{"🥇🏆", "🥈🏆", "🥉🏆"}

	rank, err := as.repository.GetUsersRank()
	if err != nil {
		return failUnexpectedOutput()
	}

	if len(rank) == 0 {
		return successOutput(OutOpt{
			Title:       "O ranking ainda está vazio... 💭",
			Description: "Os frangos ainda não submeteram treinos para o contador...",
		})
	}

	if len(rank) > 0 && len(rank) < 3 {
		return successOutput(OutOpt{
			Title:       "Opa! Perai...",
			Description: "É necessário pelo menos ter 3 pessoas pra montar um ranking...",
		})
	}

	if len(rank) == 3 {
		for i, v := range rank[:3] {
			emojiIter += fmt.Sprintf("\nTOP %d %s - %d %s", i+1, v.Nickname, v.Count, emojis[i])
		}
		return successOutput(OutOpt{
			Title:       "ranking dos mais saudáveis e marombeiros. 💪🏅",
			Description: emojiIter,
		})
	}

	for i, v := range rank[:3] {
		emojiIter += fmt.Sprintf("\nTOP %d %s - %d %s", i+1, v.Nickname, v.Count, emojis[i])
	}
	for i, v := range rank[3:] {
		restIter += fmt.Sprintf("\nTOP %d %s - %d", i+4, v.Nickname, v.Count)
	}
	return successOutput(OutOpt{
		Title:       "Ranking dos mais saudáveis e marombeiros. 💪🏅",
		Description: emojiIter + restIter,
	})
}

func (as *ActivitiesServices) ExecuteReset(i *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	myDiscord := discord.GetUserData(i)
	modsId := strings.Split(env.Getenv("MODS_ID"), ",")
	iamMod := slices.Contains(modsId, myDiscord.Id)
	if !iamMod {
		return failOutput(OutOpt{
			Title:       "🤡🤡🤡🤡🤡🤡🤡",
			Description: "🐰 Alice, curiosa como sempre, seguiu um coelho branco até um buraco misterioso. O que poderia dar errado,Alice? 🐰",
		})
	}
	if err := as.repository.ResetCount(); err != nil {
		return failUnexpectedOutput()
	}
	return successOutput(OutOpt{
		Title:       "Contagem resetada com sucesso!",
		Description: fmt.Sprintf("%s usou o comando para resetar as contagens dos frangos!", myDiscord.Nickname),
	})
}

func (as *ActivitiesServices) HelpCmd() *discordgo.InteractionResponseData {
	description := fmt.Sprintln(`
		/inscrever: Este comando te incluirá na lista de contagem de treinos o autor do comando. ✅ 

		/ta-pago: Este comando validara a contagem de treino do autor do comando, aumentando sua posição no ranking. 💪
	
		/ranking: Use este comando para visualizar a lista atualizada dos **10 Primeiros** participantes. 🏆🏅

		/reset: Este comando é utilizado pelos administradores do servidor para resetar a contagem de treinos caso algo dê problema. 🫡💪
		`)
	return successOutput(OutOpt{
		Title:       "Veja abaixo como os comandos funcionam.",
		Description: description,
	})
}
