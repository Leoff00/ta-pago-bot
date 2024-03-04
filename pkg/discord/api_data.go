package discord

import "github.com/bwmarrin/discordgo"

type UserData struct {
	Id        string `json:"id"`
	ChannelId string `json:"channel_id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Mention   string `json:"mention"`
	ServerId  string `json:"server_id"` //use as multi tenant id
}

// GetUserData retrieve the Discord Api user data from the interaction
// Put here any data that comes from discord
func GetUserData(i *discordgo.InteractionCreate) *UserData {
	discordUser := &UserData{
		Id:        i.Member.User.ID,
		ChannelId: i.ChannelID,
		Username:  i.Member.User.Username,
		Nickname:  i.Member.Nick,
		Mention:   i.Member.Mention(),
		ServerId:  i.GuildID,
	}
	return discordUser
}
