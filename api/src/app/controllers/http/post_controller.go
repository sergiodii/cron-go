package api_constroller_http

import (
	"net/http"

	"github.com/labstack/echo"
	api_services "github.com/sergiodii/cron-go/api/src/app/services"
	utils_api "github.com/sergiodii/cron-go/api/src/utils"
)

func CreatePost(ctx echo.Context) error {

	if err := api_services.PostCreateService(ctx.Request().Body); err != nil {
		utils_api.Logger.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	utils_api.Logger.Info(http.StatusCreated, map[string]string{"message": "post created"})
	return ctx.JSON(http.StatusCreated, map[string]string{"message": "post created"})
}

func GetAllPosts(ctx echo.Context) error {

	data, err := api_services.PostsGetAllService(ctx)
	if err != nil {
		utils_api.Logger.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	utils_api.Logger.Info(http.StatusCreated, data)
	return ctx.JSON(http.StatusCreated, data)
}
