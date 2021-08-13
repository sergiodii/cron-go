package respository_implementation

import (
	"fmt"

	"github.com/sergiodii/cron-go/domain/entities"
	"github.com/sergiodii/cron-go/domain/interfaces"
	"github.com/sergiodii/cron-go/domain/repositories"
	"github.com/sergiodii/cron-go/shared"
)

type postsUserPostPGRepository struct {
	logger     interface{}
	connection *shared.PGDB
}

func NewUserPostPGRepository() repositories.IUserPostRepository {
	job := postsUserPostPGRepository{}
	job.connection = shared.NewPGDB("")
	return &job
}

func (pr *postsUserPostPGRepository) GetAll(parans interfaces.UserPostQueryParans) []interfaces.UserPostResponseDTO {
	pr.connection.Connect()
	defer pr.connection.Disconnect()

	var posts []entities.PostsEntity

	orderByString := ""
	if len(parans.Order) > 0 {
		orderByString = "ORDER BY " + parans.Order + " desc"
	}
	pr.connection.Database.Raw(fmt.Sprintf("SELECT * FROM posts_entities %s", orderByString)).Scan(&posts)
	var result []interfaces.UserPostResponseDTO
	for _, p := range posts {
		var r interfaces.UserPostResponseDTO
		r.Name = p.Author
		result = append(result, r)
	}
	return result
}
