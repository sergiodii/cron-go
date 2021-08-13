package api_config

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
	utils_api "github.com/sergiodii/cron-go/api/src/utils"
	"github.com/sergiodii/cron-go/domain/entities"
	"github.com/sergiodii/cron-go/domain/use_cases"
	"github.com/sergiodii/cron-go/shared"
)

func Start() {
	envFileName := shared.GetEnvFileName()

	if err := godotenv.Load(envFileName); err != nil {
		utils_api.Logger.Warning("Shared Package dont founded ", envFileName, " file")
	}

	if err := use_cases.MigrateTablePostUseCase.Execute(); err != nil {
		utils_api.Logger.Error(err)
	}

	elasticIndexPost := GetElasticIndexPostData()

	if !VerifyIfElasticIndexIsCreated(elasticIndexPost) {
		use_cases.CreateIndexElasticUseCase.Execute(elasticIndexPost)
	}

}

func VerifyIfElasticIndexIsCreated(data entities.ElasticIndexEntity) bool {
	url := shared.GetEnvOrFail("ELASTIC_SEARCH_URL")
	resp, err := http.Get(fmt.Sprint(url, "/", strings.ToLower(data.Name)))
	if err != nil {
		utils_api.Logger.Error(err)
	}
	if resp.StatusCode < 300 && resp.StatusCode >= 200 {
		return true
	} else if resp.StatusCode >= 500 {
		return true
	}
	return false
}

func GetElasticIndexPostData() entities.ElasticIndexEntity {
	var a entities.ElasticIndexEntity
	a.Name = "posts"
	a.Mappings.Properties = map[string]map[string]string{
		"creation_date": {
			"type": "date",
		},
		"author": {
			"type": "text",
		},
		"title": {
			"type": "text",
		},
		"ups": {
			"type": "integer",
		},
		"num_comments": {
			"type": "integer",
		},
	}
	a.Settings.NumberOfReplicas = 1
	a.Settings.NumberOfShards = 1
	return a
}
