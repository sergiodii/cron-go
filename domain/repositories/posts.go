package repositories

import (
	"github.com/sergiodii/cron-go/domain/entities"
	"github.com/sergiodii/cron-go/domain/interfaces"
)

type IPostsRepository interface {
	GetAll(parans interfaces.PostQueryParans) []entities.PostsEntity
	// GetById(id uint) entities.PostsEntity
	// GetByName(name string) entities.PostsEntity
	Create(post entities.PostsEntity) (id uint, err error)
	AutoMigrate() error
}
