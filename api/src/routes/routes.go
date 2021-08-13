package routes

import (
	"net/http"

	"github.com/labstack/echo"
	api_constroller_http "github.com/sergiodii/cron-go/api/src/app/controllers/http"
)

func Routes(route *echo.Echo) {

	route.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "CRON-GO-API")
	})

	route.GET("/user-posts", api_constroller_http.GetAllUsersPosts)
	route.GET("/posts", api_constroller_http.GetAllPosts)
	route.POST("/posts", api_constroller_http.CreatePost)
}
