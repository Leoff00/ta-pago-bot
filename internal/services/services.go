package services

import (
	"errors"
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
	mentionReference := i.Member.Mention()
	discordUser := discord.GetUserData(i)
	alreadyExists, _ := as.repository.ExistsById(discordUser.Id)
	if alreadyExists {
		err := errors.New(fmt.Sprintf("Parece que o canela seca do %s ta tentando me derrubar, TU JA TA INSCRITO SUA MULA!! ", mentionReference))
		return failOutput(err)
	}
	user, err := domain.NewUser(domain.CreateUserOpts{
		Id:       discordUser.Id,
		Username: discordUser.Username,
		Nickname: discordUser.Nickname,
	})
	if err != nil {
		log.Default().Println("Error during user creation", err.Error())
		return failOutput(errors.New("Ocorreu um erro inesperado. Não foi possível criar o usuário"))

	}
	err = as.repository.Create(user)
	if err != nil {
		return failOutput(errors.New("Ocorreu um erro inesperado. Não foi possível criar o usuário"))
	}

	title := "Agora é só mandar bala, digite o comando /ta-pago toda vez que buscar o shape meu nobre!! 💪🏅"
	description := helpers.RandomizeJoinPhrases(mentionReference)
	return successOutput(title, description)
}

func (as *ActivitiesServices) ExecutePay(i *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	discordUsr := discord.GetUserData(i)
	user, err := as.repository.GetUserById(i.Member.User.ID)
	if err != nil {
		return failOutput(nil)
	}
	aggregate := &models.UserAggregate{
		User:        user,
		DiscordUser: discordUsr,
	}
	if user.IsNotSubscribed() {
		err := errors.New("você precisa antes se inscrever na lista fera")
		return failOutput(err)
	}
	if user.AlreadySubmitted() {
		err := errors.New("seu frango! tu já treinou hoje mermão, volta amanhã")
		return failOutput(err)
	}
	user.Pay()
	err = as.repository.Save(aggregate)
	if err != nil {
		return failOutput(nil)
	}
	title := fmt.Sprintf("%s pagou!!!", user.GetNickname())
	description := helpers.RandomizePayPhrases(discordUsr.Member.Mention())
	return successOutput(title, description)
}

func (as *ActivitiesServices) ExecuteRanking() *discordgo.InteractionResponseData {
	var emojiIter string
	var restIter string
	emojis := [3]string{"🥇🏆", "🥈🏆", "🥉🏆"}

	rank, err := as.repository.GetUserRank()
	if err != nil {
		return failOutput(errors.New("Erro inesperado. Verifica essa parada ai !!"))
	}

	if len(rank) == 0 {
		title := "O ranking ainda está vazio... 💭"
		description := "Os frangos ainda não submeteram treinos para o contador..."
		return successOutput(title, description)
	}

	if len(rank) > 0 && len(rank) < 3 {
		title := "Opa! Perai..."
		description := "É necessário pelo menos ter 3 pessoas pra montar um ranking..."
		return successOutput(title, description)
	}

	if len(rank) == 3 {
		for i, v := range rank[:3] {
			emojiIter += fmt.Sprintf("\nTOP %d %s - %d %s", i+1, v.Nickname, v.Count, emojis[i])
		}
		title := "ranking dos mais saudáveis e marombeiros. 💪🏅"
		description := emojiIter
		return successOutput(title, description)
	}

	for i, v := range rank[:3] {
		emojiIter += fmt.Sprintf("\nTOP %d %s - %d %s", i+1, v.Nickname, v.Count, emojis[i])
	}

	for i, v := range rank[3:] {
		restIter += fmt.Sprintf("\nTOP %d %s - %d", i+4, v.Nickname, v.Count)
	}

	finalStr := emojiIter + restIter
	title := "ranking dos mais saudáveis e marombeiros. 💪🏅"
	return successOutput(title, finalStr)
}

func (as *ActivitiesServices) ExecuteReset(i *discordgo.InteractionCreate) *discordgo.InteractionResponseData {
	myDiscord := discord.GetUserData(i)
	modsId := strings.Split(env.Getenv("MODS_ID"), ",")
	iamMod := slices.Contains(modsId, myDiscord.Id)
	if !iamMod {
		log.Default().Println(fmt.Sprintf("O usuário %s tentou resetar a contagem sem permissão", myDiscord.Id))
		title := "🤡🤡🤡🤡🤡🤡🤡"
		description := "🐰 Alice, curiosa como sempre, seguiu um coelho branco até um buraco misterioso. O que poderia dar errado,Alice? 🐰"
		return customFailOutput(title, description)
	}
	if err := as.repository.ResetCount(); err != nil {
		return failOutput(errors.New("Erro inesperado. Verifica essa parada ai mermão!!"))
	}
	fmtDescription := fmt.Sprintf("%s usou o comando para resetar as contagens dos frangos!", myDiscord.Nickname)
	return successOutput("TABELA RESETADA!!!", fmtDescription)
}

func (as *ActivitiesServices) HelpCmd() *discordgo.InteractionResponseData {
	description := fmt.Sprintln(`
		/inscrever: Este comando te incluirá na lista de contagem de treinos o autor do comando. ✅ 

		/ta-pago: Este comando validara a contagem de treino do autor do comando, aumentando sua posição no ranking. 💪
	
		/ranking: Use este comando para visualizar a lista atualizada dos **10 Primeiros** participantes. 🏆🏅

		/reset: Este comando é utilizado pelos administradores do servidor para resetar a contagem de treinos caso algo dê problema. 🫡💪
		`)
	title := "Veja abaixo como os comandos funcionam."
	return successOutput(title, description)
}

// -> revise if this is the correct place to construct these outputs
// only domain errors here (not db, userRepo, disc or sensitive errors)
func failOutput(err error) *discordgo.InteractionResponseData {
	description := "Ocorreu um erro inesperado. Não foi possível criar o usuário"
	if err != nil {
		description = err.Error()
	}
	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{
			&discordgo.MessageEmbed{
				Title:       "Deu merda aqui!!!",
				Description: fmt.Sprintln(description + "❌"),
				Type:        discordgo.EmbedTypeRich,
				Color:       10,
			}},
	}
}

func customFailOutput(title string, description string) *discordgo.InteractionResponseData {
	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{
			&discordgo.MessageEmbed{
				Title:       title,
				Description: description,
				Type:        discordgo.EmbedTypeRich,
				Color:       10,
			}},
	}
}

func successOutput(title string, description string) *discordgo.InteractionResponseData {
	return &discordgo.InteractionResponseData{
		Embeds: MsgEmbedType{
			&discordgo.MessageEmbed{
				Title:       title,
				Description: description,
				Type:        discordgo.EmbedTypeRich,
				Color:       10,
			}},
	}
}
