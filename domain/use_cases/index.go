package use_cases

import (
	"github.com/sergiodii/cron-go/domain/repositories"
	respository_implementation "github.com/sergiodii/cron-go/domain/repositories/implementation"
	use_cases_jobs "github.com/sergiodii/cron-go/domain/use_cases/jobs"
	use_cases_posts "github.com/sergiodii/cron-go/domain/use_cases/posts"
)

// =============================
// JOBS
// =============================

var jobRepository repositories.IJobsRepository = respository_implementation.NewJobsPGRepository()

var GetJobsUseCase use_cases_jobs.GetJobsUseCase = *use_cases_jobs.NewGetJobsUseCase(jobRepository)
var GetJobByIdUseCase use_cases_jobs.GetJobByIdUseCase = *use_cases_jobs.NewGetJobByIdUseCase(jobRepository)
var CreateJobUseCase use_cases_jobs.CreateJobUseCase = *use_cases_jobs.NewCreateJobUseCase(jobRepository)
var MigrateTableJobUseCase use_cases_jobs.MigrateTableJobUseCase = *use_cases_jobs.NewMigrateTableJobUseCase(jobRepository)

// =============================
// POSTS
// =============================

var postResitory repositories.IPostsRepository = respository_implementation.NewPostsPGRepository()

var CreatePostUseCase use_cases_posts.CreatePostUseCase = *use_cases_posts.NewCreatePostUseCase(postResitory)
