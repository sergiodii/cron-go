package entities

import "encoding/json"

type ElasticIndexEntity struct {
	Settings struct {
		NumberOfReplicas int64 `json:"number_of_replicas"`
		NumberOfShards   int64 `json:"number_of_shards"`
	} `json:"settings"`
	Mappings struct {
		Properties map[string]map[string]string `json:"properties"`
	} `json:"mappings"`
	Name string `json:"-"`
}

func (els *ElasticIndexEntity) FromJson(dataJson string) *ElasticIndexEntity {
	jobsent := new(ElasticIndexEntity)
	json.Unmarshal([]byte(dataJson), &jobsent)

	return jobsent
}

func (els *ElasticIndexEntity) ToJson() string {
	j, e := json.Marshal(els)
	if e != nil {
		return `{"error":"error"}`
	}
	return string(j)
}
