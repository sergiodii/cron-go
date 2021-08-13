package use_cases_jobs

import (
	"github.com/sergiodii/cron-go/domain/entities"
	"github.com/sergiodii/cron-go/domain/repositories"
)

type CreateJobUseCase struct {
	repository repositories.IJobsRepository
}

func NewCreateJobUseCase(repository repositories.IJobsRepository) *CreateJobUseCase {
	createJobUseCase := new(CreateJobUseCase)
	createJobUseCase.repository = repository
	return createJobUseCase
}

func (gjob *CreateJobUseCase) Execute(job entities.JobsEntity) (uint, error) {
	id, err := gjob.repository.Create(job)
	return id, err
}
