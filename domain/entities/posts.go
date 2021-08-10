package entities

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type PostsEntity struct {
	gorm.Model
	Id           uint       `gorm:"primaryKey" json:"id"`
	Author       string     `json:"author"`
	Title        string     `json:"title"`
	Ups          int64      `json:"ups"`
	NumComments  int64      `json:"num_comments"`
	CreationData time.Time  `json:"created_data"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

func (pe *PostsEntity) FromJson(dataJson string) *PostsEntity {
	postent := new(PostsEntity)
	json.Unmarshal([]byte(dataJson), &postent)

	return postent
}

func (pe *PostsEntity) ToJson() string {
	j, e := json.Marshal(pe)
	if e != nil {
		return `{"error":"error"}`
	}
	return string(j)
}
