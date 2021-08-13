package use_cases_posts

import (
	"github.com/sergiodii/cron-go/domain/entities"
	"github.com/sergiodii/cron-go/domain/interfaces"
	"github.com/sergiodii/cron-go/domain/repositories"
)

type GetPostUseCase struct {
	repository repositories.IPostsRepository
}

func NewGetPostUseCase(repository repositories.IPostsRepository) *GetPostUseCase {
	getPostsUseCase := new(GetPostUseCase)
	getPostsUseCase.repository = repository
	return getPostsUseCase
}

func (gjob *GetPostUseCase) Execute(parans interfaces.PostQueryParans) []entities.PostsEntity {
	jobs := gjob.repository.GetAll(parans)
	return jobs
}
