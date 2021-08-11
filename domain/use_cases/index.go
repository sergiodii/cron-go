package use_cases

import (
	"github.com/sergiodii/cron-go/domain/repositories"
	respository_implementation "github.com/sergiodii/cron-go/domain/repositories/implementation"
	use_cases_jobs "github.com/sergiodii/cron-go/domain/use_cases/cron_jobs"
	use_cases_posts "github.com/sergiodii/cron-go/domain/use_cases/posts"
	use_cases_user_posts "github.com/sergiodii/cron-go/domain/use_cases/user_posts"
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

var postRepository repositories.IPostsRepository = respository_implementation.NewPostsPGRepository()

var GetPostUseCase use_cases_posts.GetPostUseCase = *use_cases_posts.NewGetPostUseCase(postRepository)
var CreatePostUseCase use_cases_posts.CreatePostUseCase = *use_cases_posts.NewCreatePostUseCase(postRepository)
var MigrateTablePostUseCase use_cases_posts.MigrateTablePostUseCase = *use_cases_posts.NewMigrateTablePostUseCase(postRepository)

// =============================
// POSTS
// =============================

var userPostRepository repositories.IUserPostRepository = respository_implementation.NewUserPostPGRepository()
var GetUserPostUseCase use_cases_user_posts.GetUserPostUseCase = *use_cases_user_posts.NewGetUserPostUseCase(userPostRepository)
