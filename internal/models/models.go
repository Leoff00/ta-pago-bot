package models

import (
	"github.com/leoff00/ta-pago-bot/internal/domain"
	"github.com/leoff00/ta-pago-bot/pkg/discord"
)

type DiscordRankType struct {
	Nickname string `json:"nickname"`
	Count    int    `json:"count"`
}

type UserAggregate struct {
	User        *domain.User
	DiscordUser *discord.UserData
}
