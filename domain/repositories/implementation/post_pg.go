package respository_implementation

import (
	"github.com/sergiodii/cron-go/domain/entities"
	"github.com/sergiodii/cron-go/domain/repositories"
	"github.com/sergiodii/cron-go/shared"
)

type postsPGRepository struct {
	logger     interface{}
	connection *shared.PGDB
}

func NewPostsPGRepository() repositories.IPostsRepository {
	job := postsPGRepository{}
	job.connection = shared.NewPGDB("")
	return &job
}

func (pr *postsPGRepository) GetAll() []entities.PostsEntity {
	pr.connection.Connect()
	defer pr.connection.Disconnect()

	var posts []entities.PostsEntity

	pr.connection.Database.Find(&posts)

	return posts
}

func (pr *postsPGRepository) Create(post entities.PostsEntity) (id uint, err error) {
	pr.connection.Connect()
	defer pr.connection.Disconnect()

	result := pr.connection.Database.Create(&post) // pass pointer of data to Create
	id = post.ID
	err = result.Error
	return
}

func (pr *postsPGRepository) AutoMigrate() error {
	pr.connection.Connect()
	defer pr.connection.Disconnect()

	err := pr.connection.Database.AutoMigrate(&entities.PostsEntity{})
	return err
}
