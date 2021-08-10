package respository_implementation

import (
	"github.com/sergiodii/cron-go/domain/entities"
	"github.com/sergiodii/cron-go/domain/repositories"
	"github.com/sergiodii/cron-go/shared"
)

type jobsPGRepository struct {
	logger     interface{}
	connection *shared.PGDB
}

func NewJobsPGRepository() repositories.IJobsRepository {
	job := jobsPGRepository{}
	job.connection = shared.NewPGDB("")
	return &job
}

func (jbr *jobsPGRepository) GetAll() []entities.JobsEntity {
	jbr.connection.Connect()
	defer jbr.connection.Disconnect()

	var jobs []entities.JobsEntity

	jbr.connection.Database.Find(&jobs)

	return jobs
}

func (jbr *jobsPGRepository) GetById(id uint) entities.JobsEntity {
	jbr.connection.Connect()
	defer jbr.connection.Disconnect()

	var job entities.JobsEntity
	jbr.connection.Database.First(&job, id)
	return job
}

func (jbr *jobsPGRepository) GetByName(name string) entities.JobsEntity {
	jbr.connection.Connect()
	defer jbr.connection.Disconnect()

	var job entities.JobsEntity
	jbr.connection.Database.Where("name =", name).Find(&job)
	return job
}

func (jbr *jobsPGRepository) Create(job entities.JobsEntity) (id uint, err error) {
	jbr.connection.Connect()
	defer jbr.connection.Disconnect()

	result := jbr.connection.Database.Create(&job) // pass pointer of data to Create
	id = job.ID
	err = result.Error
	return
}

func (jbr *jobsPGRepository) AutoMigrate() error {
	jbr.connection.Connect()
	defer jbr.connection.Disconnect()

	err := jbr.connection.Database.AutoMigrate(&entities.JobsEntity{})
	return err
}
