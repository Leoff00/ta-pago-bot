package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/leoff00/ta-pago-bot/pkg/helpers"
)

func fmtTable() string {
	// var res string
	firstBacktick := "```markdown"
	lastBacktick := "```"

	// foo := []string{"foo", "bar", "baz"}

	// res = fmt.Sprintf(`
	// | NOMES:  | RANKING |
	// | --------| ------- |
	// | %s 			|	TOP 1   |
	// | %s 			|	TOP 2   |
	// | %s      | TOP 3   |`, foo[0], foo[1], foo[2])

	// style := fmt.Sprintf("%s\n %s \n%s", firstBacktick, res, lastBacktick)
	// fmt.Println(style)
	// return style

	return fmt.Sprintf(`
	%s
	| Month    | Savings |
	| -------- | ------- |
	| January  | $250    |
	| February | $80     |
	| March    | $420    |
	%s`, firstBacktick, lastBacktick)
}

func ExecHandlers(bot *discordgo.Session, botId string) {
	bot.AddHandlerOnce(OnReady())
	bot.AddHandler(activities(botId))
}

func activities(botId string) InteractionCreateResponse {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		msgEmbed := &discordgo.MessageEmbed{
			Title:       "Ranking dos mais saúdaveis e marombeiros",
			Description: fmtTable(),
			Type:        discordgo.EmbedTypeRich,
			Color:       10,
		}
		dt := &discordgo.InteractionResponseData{
			Embeds: MsgEmbedType{msgEmbed},
		}
		if i.Type == AppCmd {
			switch i.ApplicationCommandData().Name {
			case "inscrever":
				helpers.InteractionResponseFactory(botId, dt, s, i)
			case "ta-pago":
				helpers.InteractionResponseFactory(botId, dt, s, i)
			case "ranking":
				helpers.InteractionResponseFactory(botId, dt, s, i)

			}
		}
	}
}
