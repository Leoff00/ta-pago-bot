package discord

import "github.com/bwmarrin/discordgo"

type UserData struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Member   *discordgo.Member
}

// GetUserData retrieve the Discord Api user data from the interaction
// Put here any data that comes from discord
func GetUserData(i *discordgo.InteractionCreate) *UserData {
	discordUser := &UserData{
		Id:       i.Member.User.ID,
		Username: i.Member.User.Username,
		Nickname: i.Member.Nick,
		Member:   i.Member,
	}
	return discordUser
}
