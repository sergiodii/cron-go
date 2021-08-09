package use_cases

import (
	"github.com/sergiodii/cron-go/domain/repositories"
	implementation "github.com/sergiodii/cron-go/domain/repositories/implementation"
	use_cases_jobs "github.com/sergiodii/cron-go/domain/use_cases/jobs"
)

var jobRepository repositories.IJobsRepository = implementation.NewJobsPGRepository()

var GetJobsUseCase use_cases_jobs.GetJobsUseCase = *use_cases_jobs.NewGetJobsUseCase(jobRepository)
var GetJobByIdUseCase use_cases_jobs.GetJobByIdUseCase = *use_cases_jobs.NewGetJobByIdUseCase(jobRepository)
var CreateJobUseCase use_cases_jobs.CreateJobUseCase = *use_cases_jobs.NewCreateJobUseCase(jobRepository)
var MigrateTableJobUseCase use_cases_jobs.MigrateTableJobUseCase = *use_cases_jobs.NewMigrateTableJobUseCase(jobRepository)
