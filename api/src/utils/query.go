package utils_api

import (
	"time"

	"github.com/labstack/echo"
)

func GetOrderFromQueryString(ctx echo.Context) string {
	orderBy := ctx.QueryParam("orderby")

	if orderBy == "up" || orderBy == "ups" {
		return "ups"
	} else if orderBy == "comment" || orderBy == "comments" {
		return "num_comments"
	}
	return ""
}

func GetStartAndEndDateFromQueryString(ctx echo.Context) (startDate time.Time, endDate time.Time) {
	dateLayout := "2006-01-02"
	startDateString := ctx.QueryParam("startdate")
	endDateString := ctx.QueryParam("enddate")

	if len(startDateString) > 0 {
		d, err := time.Parse(dateLayout, startDateString)
		if err != nil {
			Logger.Error(err)
		}
		startDate = d
	}

	if len(endDateString) > 0 {
		d, err := time.Parse(dateLayout, endDateString)
		if err != nil {
			Logger.Error(err)
		}
		endDate = d
		endDate.Add(23 * time.Hour)
	}

	return
}
