package error_handle

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	utils_api "github.com/sergiodii/cron-go/api/src/utils"
)

func Handle(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	utils_api.Logger.Error(report.Message)
	c.HTML(report.Code, report.Message.(string))
}
