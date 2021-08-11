package api_constroller_http

import (
	"net/http"

	"github.com/labstack/echo"
	// api_services "github.com/sergiodii/cron-go/api/src/app/services"
	api_services "github.com/sergiodii/cron-go/api/src/app/services"
	utils_api "github.com/sergiodii/cron-go/api/src/utils"
)

func GetAllUsersPosts(ctx echo.Context) error {
	userPostsList, err := api_services.UserPostsGetAllService(ctx)
	if err != nil {
		utils_api.Logger.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	utils_api.Logger.Info(http.StatusCreated, userPostsList)
	return ctx.JSON(http.StatusSeeOther, userPostsList)
}
