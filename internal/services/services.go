package services

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/leoff00/ta-pago-bot/internal/domain"
	"github.com/leoff00/ta-pago-bot/internal/helpers"
	"github.com/leoff00/ta-pago-bot/internal/models"
	"github.com/leoff00/ta-pago-bot/internal/models/tenants"
	"github.com/leoff00/ta-pago-bot/internal/repo"
	"github.com/leoff00/ta-pago-bot/pkg/discord"
)

// ActivitiesServices is a struct that holds configs & repository from a tenant
type ActivitiesServices struct {
	tenants map[string]*tenants.Tenant
}

func NewActivitiesServices(tenants map[string]*tenants.Tenant) *ActivitiesServices {
	return &ActivitiesServices{
		tenants: tenants,
	}
}

func (as *ActivitiesServices) ExecuteJoin(
	i *discordgo.InteractionCreate,
) *discordgo.InteractionResponseData {
	discordUsr := discord.GetUserData(i)
	repository := as.GetRepository(discordUsr.ServerId)

	allowedChannelId := as.tenants[discordUsr.ServerId].ChannelID

	if discordUsr.ChannelId != allowedChannelId {
		return failOutput(OutOpt{
			Description: "Só é permitido enviar comandos do bot no canal configurado!!!",
		})
	}

	isSubscribed, err := repository.ExistsById(discordUsr.Id)
	if err != nil {
		return failUnexpectedOutput()
	}
	if isSubscribed {
		d := fmt.Sprintf(
			"Parece que o canela seca do %s ta tentando me derrubar, TU JA TA INSCRITO SUA MULA!!",
			discordUsr.Mention)
		return failOutput(OutOpt{
			Description: d,
		})
	}
	user, err := domain.NewUser(domain.CreateUserOpts{
		Id:       discordUsr.Id,
		Username: discordUsr.Username,
		Nickname: discordUsr.Nickname,
	})
	if err != nil {
		log.Default().Println("Error during user creation", err.Error())
		return failUnexpectedOutput()
	}
	if err = repository.Insert(user); err != nil {
		return failUnexpectedOutput()
	}
	return successOutput(OutOpt{
		Title:       "Agora é só mandar bala, digite o comando /ta-pago toda vez que buscar o shape meu nobre!! 💪🏅",
		Description: helpers.RandomizeJoinPhrases(discordUsr.Mention),
	})
}

func (as *ActivitiesServices) ExecutePay(
	i *discordgo.InteractionCreate,
) *discordgo.InteractionResponseData {
	discordUsr := discord.GetUserData(i)
	repository := as.GetRepository(discordUsr.ServerId)

	user, err := repository.GetUserById(i.Member.User.ID)
	allowedChannelId := as.tenants[discordUsr.ServerId].ChannelID

	if discordUsr.ChannelId != allowedChannelId {
		return failOutput(OutOpt{
			Description: "Só é permitido enviar comandos do bot no canal configurado!!!",
		})
	}

	if err != nil {
		return failUnexpectedOutput()
	}

	aggregate := &models.UserAggregate{
		User:        user,
		DiscordUser: discordUsr,
	}

	if user.IsNotSubscribed() {
		return failOutput(OutOpt{
			Description: "você precisa antes se inscrever na lista fera",
		})
	}
	if user.AlreadySubmitted() {
		return failOutput(OutOpt{
			Description: "seu frango! tu já treinou hoje mermão, volta amanhã",
		})
	}

	user.Pay()
	err = repository.Save(aggregate)
	if err != nil {
		return failUnexpectedOutput()
	}
	return successOutput(OutOpt{
		Title:       fmt.Sprintf("%s pagou!!!", user.GetNickname()),
		Description: helpers.RandomizePayPhrases(discordUsr.Mention),
	})
}

func (as *ActivitiesServices) ExecuteRanking(
	i *discordgo.InteractionCreate,
	serverId string,
) *discordgo.InteractionResponseData {
	discordUsr := discord.GetUserData(i)
	var emojiIter string
	var restIter string
	emojis := [3]string{"🥇🏆", "🥈🏆", "🥉🏆"}

	if serverId == "" {
		serverId = discord.GetUserData(i).ServerId
	}

	allowedChannelId := as.tenants[discordUsr.ServerId].ChannelID

	if discordUsr.ChannelId != allowedChannelId {
		return failOutput(OutOpt{
			Description: "Só é permitido enviar comandos do bot no canal configurado!!!",
		})
	}

	repository := as.GetRepository(serverId)

	rank, err := repository.GetUsersRank()
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

func (as *ActivitiesServices) ExecuteReset(
	i *discordgo.InteractionCreate,
) *discordgo.InteractionResponseData {
	discordUsr := discord.GetUserData(i)
	modsId := as.tenants[discordUsr.ServerId].ModsID
	iamMod := slices.Contains(modsId, discordUsr.Id)

	allowedChannelId := as.tenants[discordUsr.ServerId].ChannelID

	if discordUsr.ChannelId != allowedChannelId {
		return failOutput(OutOpt{
			Description: "Só é permitido enviar comandos do bot no canal configurado!!!",
		})
	}

	if !iamMod {
		return failOutput(OutOpt{
			Title:       "🤡🤡🤡🤡🤡🤡🤡",
			Description: "🐰 Alice, curiosa como sempre, seguiu um coelho branco até um buraco misterioso. O que poderia dar errado,Alice? 🐰",
		})
	}
	repository := as.GetRepository(discordUsr.ServerId)
	if err := repository.ResetCount(); err != nil {
		return failUnexpectedOutput()
	}
	return successOutput(OutOpt{
		Title: "Contagem resetada com sucesso!",
		Description: fmt.Sprintf(
			"%s usou o comando para resetar a contagem dos frangos!",
			discordUsr.Username,
		),
	})
}

func (as *ActivitiesServices) ExecuteModEditCount(
	i *discordgo.InteractionCreate,
) *discordgo.InteractionResponseData {
	// max value user can do workout (measured in months)
	const MAX_WORKOUT_COUNT_ALLOWED int8 = 31
	discordUsr := discord.GetUserData(i)
	userId := i.ApplicationCommandData().Options[0].StringValue()
	countValue := i.ApplicationCommandData().Options[1].IntValue()

	allowedChannelId := as.tenants[discordUsr.ServerId].ChannelID

	if discordUsr.ChannelId != allowedChannelId {
		return failOutput(OutOpt{
			Description: "Só é permitido enviar comandos do bot no canal configurado!!!",
		})
	}

	normalizedUserId := strings.Join(regexp.MustCompile(`\d`).FindAllString(userId, -1), "")

	if int8(countValue) < 0 || int8(countValue) > MAX_WORKOUT_COUNT_ALLOWED {
		return failOutput(OutOpt{
			Title:       "Deu Merda aqui!!!",
			Description: "So é permitido editar a contagem até 31 meu nobre!!!",
		})
	}

	repository := as.GetRepository(discordUsr.ServerId)

	modsId := as.tenants[discordUsr.ServerId].ModsID
	iamMod := slices.Contains(modsId, discordUsr.Id)

	if !iamMod {
		return failOutput(OutOpt{
			Title:       "🤡🤡🤡🤡🤡🤡🤡",
			Description: "🐰 Alice, curiosa como sempre, seguiu um coelho branco até um buraco misterioso. O que poderia dar errado,Alice? 🐰",
		})
	}

	if err := repository.EditCount(normalizedUserId, int(countValue)); err != nil {
		failUnexpectedOutput()
	}

	return successOutput(OutOpt{
		Title: "Contagem editada!",
		Description: fmt.Sprintf(
			"%s Usou o comando para editar a contagem dos frangos!",
			discordUsr.Mention,
		),
	})
}

func (as *ActivitiesServices) HelpCmd() *discordgo.InteractionResponseData {
	description := fmt.Sprintln(`
		/inscrever: Este comando te incluirá na lista de contagem de treinos o autor do comando. ✅ 

		/ta-pago: Este comando validara a contagem de treino do autor do comando, aumentando sua posição no ranking. 💪
	
		/ranking: Use este comando para visualizar a lista atualizada dos **10 Primeiros** participantes. 🏆🏅

		/editar_contagem: Este comando é utilizado pelos administradores do servidor para editar a contagem caso ocorra adversidades. 💪

		/reset: Este comando é utilizado pelos administradores do servidor para resetar a contagem de treinos caso algo dê problema. 🫡💪
		`)
	return successOutput(OutOpt{
		Title:       "Veja abaixo como os comandos funcionam.",
		Description: description,
	})
}

func (as *ActivitiesServices) GetRepository(serverID string) *repo.UserRepository {
	return as.tenants[serverID].Repository
}
