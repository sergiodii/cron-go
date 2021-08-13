package use_cases_elastic

import (
	"github.com/sergiodii/cron-go/domain/repositories"
)

type DataInsertionIndexElasticUseCase struct {
	repository repositories.IElasticSearchRepository
}

func NewDataInsertionIndexElasticUseCase(repository repositories.IElasticSearchRepository) *DataInsertionIndexElasticUseCase {
	dataInsertionIndexElasticSearch := new(DataInsertionIndexElasticUseCase)
	dataInsertionIndexElasticSearch.repository = repository
	return dataInsertionIndexElasticSearch
}

func (els *DataInsertionIndexElasticUseCase) Execute(index string, model interface{}) error {
	return els.repository.DataInsertion(index, model)
}
