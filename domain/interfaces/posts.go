package interfaces

import "time"

type PostQueryParans struct {
	Order     string
	StartDate time.Time
	EndDate   time.Time
}

type ElasticSearchPostModel struct {
	CreationDate time.Time `json:"creation_date"`
	Author       string    `json:"author"`
	Title        string    `json:"title"`
	Ups          int       `json:"ups"`
	NumComments  int       `json:"num_comments"`
}
