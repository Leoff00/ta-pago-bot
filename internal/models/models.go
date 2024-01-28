package models

type DiscordUser struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Count      int    `json:"count"`
	Updated_at int    `json:"updated_at"`
}

type DiscordReturnType struct {
	Username string `json:"username"`
	Count    int    `json:"count"`
}
