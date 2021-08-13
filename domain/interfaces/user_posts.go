package interfaces

import "time"

type Order string

type UserPostQueryParans struct {
	Order     string
	StartDate time.Time
	EndDate   time.Time
}

type UserPostResponseDTO struct {
	Name string
}
