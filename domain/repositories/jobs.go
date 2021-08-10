package repositories

import "github.com/sergiodii/cron-go/domain/entities"

type IJobsRepository interface {
	GetAll() []entities.JobsEntity
	GetById(id uint) entities.JobsEntity
	GetByName(name string) entities.JobsEntity
	Create(job entities.JobsEntity) (id uint, err error)
	AutoMigrate() error
}
