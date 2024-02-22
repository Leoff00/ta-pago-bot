package tenants

import "github.com/leoff00/ta-pago-bot/internal/repo"

// Tenant struct contains data that every discord server needs to have
// specific folder to avoid circular imports
type Tenant struct {
	ServerName string   `json:"server_name"`
	ServerID   string   `json:"server_id"`
	ChannelID  string   `json:"channel_id"`
	ModsID     []string `json:"mods_id"`
	DBName     string   `json:"db_name"`
	Repository *repo.UserRepository
}
