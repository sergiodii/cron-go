package api_config

import (
	"github.com/joho/godotenv"
	utils_api "github.com/sergiodii/cron-go/api/src/utils"
	"github.com/sergiodii/cron-go/domain/use_cases"
	"github.com/sergiodii/cron-go/shared"
)

func Start() {
	envFileName := shared.GetEnvFileName()

	if err := godotenv.Load(envFileName); err != nil {
		utils_api.Logger.Warning("Shared Package dont founded ", envFileName, " file")
	}

	if err := use_cases.MigrateTablePostUseCase.Execute(); err != nil {
		utils_api.Logger.Error(err)
	}
}
