package helpers

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/leoff00/ta-pago-bot/internal/repo"
	"github.com/leoff00/ta-pago-bot/internal/services"
	"github.com/leoff00/ta-pago-bot/pkg/env"
	"github.com/robfig/cron/v3"
)

var (
	dur = repo.DiscordUserRepository{}
)

type CronTasks struct {
	Cron *cron.Cron
}

func (ct *CronTasks) ScheduleWipeCountMessage(s *discordgo.Session) {
	c := ct.Cron

	if _, err := c.AddFunc("@monthly", func() {

		if err := dur.RestartCount(); err != nil {
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
	if _, err := c.AddFunc("0 12 * * FRI", func() {
		as := services.ActivitiesServices{}
		_, embed := as.ExecuteRankingService()
		if _, err := s.ChannelMessageSendEmbed(env.Getenv("CHANNEL_ID"), embed); err != nil {
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

func (ct *CronTasks) ExecuteTasks(s *discordgo.Session) {
	ct.ScheduleTrainMessage(s)
	ct.ScheduleRankingMessage(s)
	ct.ScheduleWipeCountMessage(s)
}
