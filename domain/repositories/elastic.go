package repositories

import "github.com/sergiodii/cron-go/domain/entities"

type IElasticSearchRepository interface {
	CreateIndex(index entities.ElasticIndexEntity)
	DataInsertion(index string, model interface{}) error
}
