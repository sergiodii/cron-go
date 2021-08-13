package respository_implementation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/sergiodii/cron-go/domain/entities"
	"github.com/sergiodii/cron-go/domain/repositories"
	"github.com/sergiodii/cron-go/shared"
)

type ElasticLocalRepository struct {
}

func NewElasticLocalRepository() repositories.IElasticSearchRepository {
	els := ElasticLocalRepository{}
	return &els
}

func (els *ElasticLocalRepository) CreateIndex(index entities.ElasticIndexEntity) {
	url := shared.GetEnvOrFail("ELASTIC_SEARCH_URL")
	logger := shared.NewLogger("")
	logger.Info("ELASTIC INDEX CREATION: REQUEST: ", url, "/", strings.ToLower(index.Name), strings.NewReader(index.ToJson()))
	req, err := http.NewRequest(http.MethodPut, fmt.Sprint(url, "/", strings.ToLower(index.Name)), strings.NewReader(index.ToJson()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.ContentLength = int64(len(index.ToJson()))
	clientHttp := &http.Client{}
	resp, err := clientHttp.Do(req)
	logElasticMessage := "ELASTIC INDEX CREATION: "
	if err != nil {
		logger.Error(logElasticMessage, err)
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(logElasticMessage, err)
	}
	logger.Info(logElasticMessage, "The calculated length is:", len(string(contents)), "for the url:", url)
	logger.Info(logElasticMessage, "   ", resp.StatusCode)
	hdr := resp.Header
	for key, value := range hdr {
		logger.Info(logElasticMessage, "   ", key, ":", value)
	}
	logger.Info(logElasticMessage, string(contents))
}

func (els *ElasticLocalRepository) DataInsertion(index string, model interface{}) error {
	logger := shared.NewLogger("")
	elasticClient, err := getElasticClient()
	if err != nil {
		logger.Error(err)
		return err
	}
	jsonBody, err := json.Marshal(model)
	if err != nil {
		logger.Error(err)
		return err
	}

	res, err := elasticClient.Index(
		index,
		strings.NewReader(string(jsonBody)),
	)
	if err != nil {
		fmt.Println("erro 3")
		logger.Error(err)
		return err
	}
	logger.Info("ELASTIC DATAINSERTION", res.StatusCode, " IS ERROR: ", res.IsError())
	return nil
}

func getElasticClient() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			shared.GetEnvOrFail("ELASTIC_SEARCH_URL"),
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return es, nil
}
