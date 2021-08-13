package middleware_api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"

	"github.com/labstack/echo/v4"
	utils_api "github.com/sergiodii/cron-go/api/src/utils"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return loggerMiddleware
}

func loggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		body := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&body)
		if err != nil {
			utils_api.Logger.Error(err)
			return err
		}

		payload := map[string]interface{}{
			"Header": utils_api.ToJsonHelper(c.Request().Header),
			"Body":   body,
			"URL":    utils_api.ToJsonHelper(c.Request().URL),
		}

		utils_api.Logger.Info(utils_api.ToJsonHelper(payload))
		return next(c)
	}
}

func ConvertIoReaderToString(data io.Reader) string {
	bodyBytes, err := ioutil.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	return bodyString
}
