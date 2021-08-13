package api_services

import (
	"encoding/json"
	"io"
	"time"

	"github.com/labstack/echo"
	utils_api "github.com/sergiodii/cron-go/api/src/utils"
	"github.com/sergiodii/cron-go/domain/entities"
	"github.com/sergiodii/cron-go/domain/interfaces"
	"github.com/sergiodii/cron-go/domain/use_cases"
	use_cases_posts "github.com/sergiodii/cron-go/domain/use_cases/posts"
)

func PostCreateService(body io.ReadCloser) error {

	var jsonData use_cases_posts.CreatePostDTO

	if err := json.NewDecoder(body).Decode(&jsonData); err != nil {
		utils_api.Logger.Error(err)
		return err
	}

	_, err := use_cases.CreatePostUseCase.Execute(jsonData)
	if err != nil {
		utils_api.Logger.Error(err)
		return err
	}

	var payload interfaces.ElasticSearchPostModel
	payload.Author = jsonData.Author
	payload.Title = jsonData.Title
	payload.CreationDate = time.Unix(jsonData.CreationDate.ToInt(), 0)
	payload.Ups = jsonData.Ups
	payload.NumComments = jsonData.NumComments

	// if err := use_cases.DataInsertionIndexElasticUseCase.Execute(api_config.GetElasticIndexPostData().Name, payload); err != nil {
	// 	utils_api.Logger.Error(err)
	// 	return err
	// }

	return nil
}

func PostsGetAllService(ctx echo.Context) ([]entities.PostsEntity, error) {

	var parans interfaces.PostQueryParans //PostQueryParans
	parans.Order = utils_api.GetOrderFromQueryString(ctx)
	startDate, endDate := utils_api.GetStartAndEndDateFromQueryString(ctx)
	parans.StartDate = startDate
	parans.EndDate = endDate
	result := use_cases.GetPostUseCase.Execute(parans)
	return result, nil
}
