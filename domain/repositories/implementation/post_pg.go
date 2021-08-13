package respository_implementation

import (
	"fmt"
	"time"

	"github.com/sergiodii/cron-go/domain/entities"
	"github.com/sergiodii/cron-go/domain/interfaces"
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

func (pr *postsPGRepository) GetAll(parans interfaces.PostQueryParans) []entities.PostsEntity {
	pr.connection.Connect()
	defer pr.connection.Disconnect()

	var posts []entities.PostsEntity

	orderByString := ""
	if len(parans.Order) > 0 {
		orderByString = "ORDER BY " + parans.Order + " desc"
	}

	setDate := false
	startDate := time.Now().AddDate(0, 0, -1)
	endDate := time.Now()

	if !parans.StartDate.IsZero() {
		setDate = true
		startDate = parans.StartDate
	}
	if !parans.EndDate.IsZero() {
		setDate = true
		endDate = parans.EndDate
	}
	dataString := ""

	if setDate {
		dataString = fmt.Sprint("WHERE creation_date >= '", startDate.Format(time.RFC3339Nano), "' AND creation_date <= '", endDate.Format(time.RFC3339Nano), "'")
	}

	pr.connection.Database.Raw(fmt.Sprintf("SELECT * FROM posts_entities %s %s", dataString, orderByString)).Scan(&posts)

	return posts
}

func (pr *postsPGRepository) Create(post entities.PostsEntity) (id uint, err error) {
	pr.connection.Connect()
	defer pr.connection.Disconnect()

	result := pr.connection.Database.Create(&post)
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
