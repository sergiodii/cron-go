package use_cases_jobs

import (
	"github.com/sergiodii/cron-go/domain/entities"
	"github.com/sergiodii/cron-go/domain/repositories"
)

type GetJobByIdUseCase struct {
	repository repositories.IJobsRepository
}

func NewGetJobByIdUseCase(repository repositories.IJobsRepository) *GetJobByIdUseCase {
	getJobByIdUseCase := new(GetJobByIdUseCase)
	getJobByIdUseCase.repository = repository
	return getJobByIdUseCase
}

func (gjob *GetJobByIdUseCase) Execute(id uint) entities.JobsEntity {
	job := gjob.repository.GetById(id)
	return job
}
