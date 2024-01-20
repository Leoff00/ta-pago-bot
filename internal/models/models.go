package models

type DiscordUser struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Count    int    `json:"count"`
}
