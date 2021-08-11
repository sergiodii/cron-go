package use_cases_posts

import (
	"github.com/sergiodii/cron-go/domain/repositories"
)

type MigrateTablePostUseCase struct {
	repository repositories.IPostsRepository
}

func NewMigrateTablePostUseCase(repository repositories.IPostsRepository) *MigrateTablePostUseCase {
	migrateJobTablePost := new(MigrateTablePostUseCase)
	migrateJobTablePost.repository = repository
	return migrateJobTablePost
}

func (gjob *MigrateTablePostUseCase) Execute() error {
	err := gjob.repository.AutoMigrate()
	return err
}
