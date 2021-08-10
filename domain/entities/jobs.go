package entities

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type JobsEntity struct {
	gorm.Model
	Id         uint           `gorm:"primaryKey" json:"-"`
	Name       string         `json:"name"`
	Cron       string         `json:"cron"`
	Min_thread uint8          `json:"min_thread"`
	Max_thread uint8          `json:"max_thread"`
	Function   string         `json:"function"`
	Parans     pq.StringArray `gorm:"type:text[]" json:"parans"`
	Runnig     bool           `json:"runnig"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  *time.Time     `json:"deleted_at"`
}

func (je *JobsEntity) FromJson(dataJson string) *JobsEntity {
	jobsent := new(JobsEntity)
	json.Unmarshal([]byte(dataJson), &jobsent)

	return jobsent
}

func (je *JobsEntity) ToJson() string {
	j, e := json.Marshal(je)
	if e != nil {
		return `{"error":"error"}`
	}
	return string(j)
}
