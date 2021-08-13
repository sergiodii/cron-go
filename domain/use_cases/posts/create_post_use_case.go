package use_cases_posts

import (
	"time"

	"github.com/sergiodii/cron-go/domain/entities"
	"github.com/sergiodii/cron-go/domain/repositories"
	"github.com/sergiodii/cron-go/shared"
)

type CreatePostDTO struct {
	Author       string           `json:"author"`
	Title        string           `json:"title"`
	Ups          int              `json:"ups"`
	NumComments  int              `json:"num_comments"`
	CreationDate shared.TimesTamp `json:"creation_date"`
}

type CreatePostUseCase struct {
	repository repositories.IPostsRepository
}

func NewCreatePostUseCase(repository repositories.IPostsRepository) *CreatePostUseCase {
	createJobUseCase := new(CreatePostUseCase)
	createJobUseCase.repository = repository
	return createJobUseCase
}

func (gjob *CreatePostUseCase) Execute(postDTO CreatePostDTO) (uint, error) {

	var post entities.PostsEntity

	post.Author = postDTO.Author
	post.Title = postDTO.Title
	post.Ups = postDTO.Ups
	post.NumComments = postDTO.NumComments
	post.CreationDate = time.Unix(postDTO.CreationDate.ToInt(), 0)

	id, err := gjob.repository.Create(post)
	return id, err
}
