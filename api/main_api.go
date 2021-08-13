package main_api

import (
	"github.com/labstack/echo"

	api_config "github.com/sergiodii/cron-go/api/src/app/config"
	"github.com/sergiodii/cron-go/api/src/app/error_handle"
	"github.com/sergiodii/cron-go/api/src/routes"
	utils_api "github.com/sergiodii/cron-go/api/src/utils"
	"go.elastic.co/apm/module/apmecho"
)

func main() {
	Main()
}

func Main() {

	api_config.Start()

	port := utils_api.GetServerPortHelper()

	server := echo.New()

	server.Use(apmecho.Middleware())

	server.HTTPErrorHandler = error_handle.Handle

	routes.Routes(server)

	server.HideBanner = true
	utils_api.Logger.Fatal(server.Start(port))
}
