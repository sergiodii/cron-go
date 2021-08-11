package interfaces

import "time"

type PostQueryParans struct {
	Order     string
	StartDate time.Time
	EndDate   time.Time
}
