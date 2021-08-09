package cron_config

import (
	cron_services "github.com/sergiodii/cron-go/cron/src/services"
	"github.com/sergiodii/cron-go/domain/use_cases"
)

func StartConfig() {
	//Migrate Job table if not exits
	use_cases.MigrateTableJobUseCase.Execute()
	cron_services.SyncJobs()

}
