package use_cases_elastic

import (
	"github.com/sergiodii/cron-go/domain/entities"
	"github.com/sergiodii/cron-go/domain/repositories"
)

type CreateIndexElasticUseCase struct {
	repository repositories.IElasticSearchRepository
}

func NewCreateIndexElasticUseCase(repository repositories.IElasticSearchRepository) *CreateIndexElasticUseCase {
	createIndexElasticSearch := new(CreateIndexElasticUseCase)
	createIndexElasticSearch.repository = repository
	return createIndexElasticSearch
}

func (els *CreateIndexElasticUseCase) Execute(data entities.ElasticIndexEntity) {
	els.repository.CreateIndex(data)
}
