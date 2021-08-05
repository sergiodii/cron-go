package middleware_api

import (
	"github.com/labstack/echo/v4"

	utils_api "github.com/sergiodii/cron-go/api/src/utils"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return loggerMiddleware
}

func loggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		payload := map[string]string{
			"Header": utils_api.ToJsonHelper(c.Request().Header),
			"Body":   utils_api.ToJsonHelper(c.Request().Body),
			"URL":    utils_api.ToJsonHelper(c.Request().URL),
		}

		utils_api.Logger.Info(utils_api.ToJsonHelper(payload))
		return next(c)
	}
}
