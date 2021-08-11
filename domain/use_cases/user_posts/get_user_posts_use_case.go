package use_cases_user_posts

import (
	"github.com/sergiodii/cron-go/domain/interfaces"
	"github.com/sergiodii/cron-go/domain/repositories"
)

type GetUserPostUseCase struct {
	repository repositories.IUserPostRepository
}

func NewGetUserPostUseCase(repository repositories.IUserPostRepository) *GetUserPostUseCase {
	getUserPostsUseCase := new(GetUserPostUseCase)
	getUserPostsUseCase.repository = repository
	return getUserPostsUseCase
}

func (gjob *GetUserPostUseCase) Execute(parans interfaces.UserPostQueryParans) []interfaces.UserPostResponseDTO {
	jobs := gjob.repository.GetAll(parans)
	return jobs
}
