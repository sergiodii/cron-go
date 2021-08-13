package repositories

import (
	"github.com/sergiodii/cron-go/domain/interfaces"
)

type IUserPostRepository interface {
	GetAll(parans interfaces.UserPostQueryParans) []interfaces.UserPostResponseDTO
}
