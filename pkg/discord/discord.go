package discord

import "github.com/bwmarrin/discordgo"

type UserData struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

// put here all necessary data that comes from discord api
func GetUserData(i *discordgo.InteractionCreate) *UserData {
	discordUser := &UserData{
		Id:       i.Member.User.ID,
		Username: i.Member.User.Username,
		Nickname: i.Member.Nick,
	}
	return discordUser
}
