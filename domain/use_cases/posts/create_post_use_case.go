package use_cases_posts

import (
	"github.com/sergiodii/cron-go/domain/entities"
	"github.com/sergiodii/cron-go/domain/repositories"
)

type CreatePostUseCase struct {
	repository repositories.IPostsRepository
}

func NewCreatePostUseCase(repository repositories.IPostsRepository) *CreatePostUseCase {
	createJobUseCase := new(CreatePostUseCase)
	createJobUseCase.repository = repository
	return createJobUseCase
}

func (gjob *CreatePostUseCase) Execute(job entities.PostsEntity) (uint, error) {
	id, err := gjob.repository.Create(job)
	return id, err
}
