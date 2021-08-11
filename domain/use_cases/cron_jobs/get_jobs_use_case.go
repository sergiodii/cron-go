package use_cases_jobs

import (
	"github.com/sergiodii/cron-go/domain/entities"
	"github.com/sergiodii/cron-go/domain/repositories"
)

type GetJobsUseCase struct {
	repository repositories.IJobsRepository
}

func NewGetJobsUseCase(repository repositories.IJobsRepository) *GetJobsUseCase {
	getJobsUseCase := new(GetJobsUseCase)
	getJobsUseCase.repository = repository
	return getJobsUseCase
}

func (gjob *GetJobsUseCase) Execute() []entities.JobsEntity {
	jobs := gjob.repository.GetAll()
	return jobs
}
