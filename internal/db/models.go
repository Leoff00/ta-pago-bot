package db

type DiscordUser struct {
	Nickname string `json:"nickname"`
	ID       string `json:"id"`
	Count    int    `json:"count"`
}
