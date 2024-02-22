package setup

import (
	"github.com/leoff00/ta-pago-bot/internal/models/tenants"
	"github.com/leoff00/ta-pago-bot/internal/repo"
	"github.com/leoff00/ta-pago-bot/internal/services"
	"github.com/leoff00/ta-pago-bot/pkg/database"
	"log"
)

func Service(tenantsCfg []tenants.Tenant, env string) *services.ActivitiesServices {
	tablesToCheck := []string{"DISCORD_USERS"}
	tenantsMap := make(map[string]*tenants.Tenant)

	// Create a map with server_id as key -> tenant as value
	for _, config := range tenantsCfg {
		db := database.SetupSqlite(config.DBName, tablesToCheck, env)
		repository := repo.NewUserRepository(db)
		config.Repository = repository
		currentConfig := config
		tenantsMap[config.ServerID] = &currentConfig
	}
	service := services.NewActivitiesServices(tenantsMap)
	logTenantSuccess(tenantsMap)
	return service

}

func logTenantSuccess(tenants map[string]*tenants.Tenant) {
	for serverID, tenant := range tenants {
		log.Default().Printf("Sucessfully setup tenants for Server %s ID %s:\n", tenant.ServerName, serverID)
	}
}
