package main_api

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	middleware_api "github.com/sergiodii/cron-go/api/src/middlewares"
	utils_api "github.com/sergiodii/cron-go/api/src/utils"
)

func main() {
	Main()
}

func Main() {

	err := godotenv.Load(".env")
	if err != nil {
		utils_api.Logger.Warning(".ENV file dont founded")
	}

	port := utils_api.GetServerPortHelper()
	server := echo.New()

	server.HideBanner = true

	server.Use(middleware_api.LoggerMiddleware())

	server.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World!")
	})

	server.Logger.Fatal(server.Start(port))
}
