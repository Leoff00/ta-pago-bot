package setup

import (
	"encoding/json"
	"log"
	"os"

	"github.com/leoff00/ta-pago-bot/internal/models/tenants"
)

func Tenants() []tenants.Tenant {
	configFile, err := os.Open("./db/tenant.json")
	if err != nil {
		log.Default().Fatalln("Can't open ./db/tenant.json", err.Error())
	}

	defer configFile.Close()

	var tenantsCfg []tenants.Tenant
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&tenantsCfg); err != nil {
		log.Default().Fatalln("Cant parse ./db/tenant.json:", err.Error())
	}
	logTenantCfgLoaded(tenantsCfg)
	return tenantsCfg
}

func logTenantCfgLoaded(tenantsCfg []tenants.Tenant) {
	for _, config := range tenantsCfg {
		log.Default().Printf(`
									Configuration for server ID: %s loaded:
									Server Name: %s
									Channel ID: %s
									Mods ID: %s
									DB Name: %s`, config.ServerID, config.ServerName, config.ChannelID, config.ModsID, config.DBName)
	}
}
