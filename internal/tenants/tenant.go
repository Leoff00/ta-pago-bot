package tenants

import (
	"github.com/leoff00/ta-pago-bot/internal/repo"
)

type Tenant struct {
	ServerName string   `json:"server_name"`
	ServerID   string   `json:"server_id"`
	ChannelID  string   `json:"channel_id"`
	ModsID     []string `json:"mods_id"`
	DBName     string   `json:"db_name"`
	Repository *repo.UserRepository
}
