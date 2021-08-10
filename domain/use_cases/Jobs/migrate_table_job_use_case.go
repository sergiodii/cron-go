package use_cases_jobs

import (
	"github.com/sergiodii/cron-go/domain/repositories"
)

type MigrateTableJobUseCase struct {
	repository repositories.IJobsRepository
}

func NewMigrateTableJobUseCase(repository repositories.IJobsRepository) *MigrateTableJobUseCase {
	migrateJobTableJob := new(MigrateTableJobUseCase)
	migrateJobTableJob.repository = repository
	return migrateJobTableJob
}

func (gjob *MigrateTableJobUseCase) Execute() error {
	err := gjob.repository.AutoMigrate()
	return err
}
