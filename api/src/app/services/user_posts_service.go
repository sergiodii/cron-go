package api_services

import (
	"github.com/labstack/echo"
	utils_api "github.com/sergiodii/cron-go/api/src/utils"
	"github.com/sergiodii/cron-go/domain/interfaces"
	"github.com/sergiodii/cron-go/domain/use_cases"
)

func UserPostsGetAllService(ctx echo.Context) ([]interfaces.UserPostResponseDTO, error) {

	var parans interfaces.UserPostQueryParans
	parans.Order = utils_api.GetOrderFromQueryString(ctx)

	result := use_cases.GetUserPostUseCase.Execute(parans)
	return result, nil
}
