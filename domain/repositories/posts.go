package repositories

import "github.com/sergiodii/cron-go/domain/entities"

type IPostsRepository interface {
	GetAll() []entities.PostsEntity
	// GetById(id uint) entities.PostsEntity
	// GetByName(name string) entities.PostsEntity
	Create(post entities.PostsEntity) (id uint, err error)
	AutoMigrate() error
}
