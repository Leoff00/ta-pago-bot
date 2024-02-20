package services

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/leoff00/ta-pago-bot/internal/repo"
	"github.com/leoff00/ta-pago-bot/pkg/env"
	"github.com/robfig/cron/v3"
)

type CronTasks struct {
	Cron     *cron.Cron
	userRepo *repo.UserRepository
	as       *ActivitiesServices
}

func NewCronService(userRepo *repo.UserRepository, as *ActivitiesServices) *CronTasks {
	return &CronTasks{
		Cron:     cron.New(),
		userRepo: userRepo,
		as:       as,
	}
}

func (ct *CronTasks) ScheduleWipeCountMessage(s *discordgo.Session) {
	c := ct.Cron
	userRepo := ct.userRepo
	if _, err := c.AddFunc("@monthly", func() {

		if err := userRepo.ResetCount(); err != nil {
			log.Default().Println(err.Error())
		}

		if _, err := s.ChannelMessageSendEmbed(env.Getenv("CHANNEL_ID"), &discordgo.MessageEmbed{
			Title:       "**CONTAGEM RESETADA**",
			Description: "@here A contagem de treinos pagos foi resetada, vá buscar seu shape e ser o primeiro!!!",
			Type:        discordgo.EmbedTypeRich,
			Color:       20,
		},
		); err != nil {
			log.Default().Println("Cannot send the message on Helper", err.Error())
		}
	}); err != nil {
		log.Default().Println("Failed to execute cron on Helper", err.Error())
	}
}

func (ct *CronTasks) ScheduleRankingMessage(s *discordgo.Session) {
	c := ct.Cron
	as := ct.as
	if _, err := c.AddFunc("0 12 * * FRI", func() {
		output := as.ExecuteRanking()
		if _, err := s.ChannelMessageSendEmbed(env.Getenv("CHANNEL_ID"), output.Embeds[0]); err != nil {
			log.Default().Println("Cannot send the message on Helper", err.Error())
		}
	}); err != nil {
		log.Default().Println("Failed to execute cron on Helper", err.Error())
	}
}

func (ct *CronTasks) ScheduleTrainMessage(s *discordgo.Session) {
	c := ct.Cron

	if _, err := c.AddFunc("@daily", func() {
		if _, err := s.ChannelMessageSendEmbed(env.Getenv("CHANNEL_ID"), &discordgo.MessageEmbed{
			Title:       "**JA TREINOU HOJE?**",
			Description: "Não deixe pra ultima hora!!! VÁ ATRAS DO SHAPE!!!",
			Type:        discordgo.EmbedTypeRich,
			Color:       20,
		},
		); err != nil {
			log.Default().Println("Cannot send the message on Helper", err.Error())
		}
	}); err != nil {
		log.Default().Println("Failed to execute cron on Helper", err.Error())
	}
}

func (ct *CronTasks) ScheduleTasks(s *discordgo.Session) {
	ct.ScheduleTrainMessage(s)
	ct.ScheduleRankingMessage(s)
	ct.ScheduleWipeCountMessage(s)
}
