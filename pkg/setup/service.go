package setup

import (
	"github.com/leoff00/ta-pago-bot/internal/repo"
	"github.com/leoff00/ta-pago-bot/internal/services"
	"github.com/leoff00/ta-pago-bot/internal/tenants"
	"log"
)

func Service(tenantsCfg []tenants.Tenant) *services.ActivitiesServices {

	// Create a map with server_id as key
	tenantsMap := make(map[string]*tenants.Tenant)
	for _, config := range tenantsCfg { //db connect
		db := DB(config.DBName)
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
