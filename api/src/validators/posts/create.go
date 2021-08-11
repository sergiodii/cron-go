package validators_api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/labstack/echo/v4"
	utils_api "github.com/sergiodii/cron-go/api/src/utils"
)

func PostsCreateValidator(c echo.Context) bool {
	payload := map[string]string{
		"Header": utils_api.ToJsonHelper(c.Request().Header),
		"Body":   ConvertIoReaderToString(c.Request().Body),
		"URL":    utils_api.ToJsonHelper(c.Request().URL),
	}

	fmt.Println("SERGIO", ConvertIoReaderToString(c.Request().Body))

	var body map[string]interface{}
	err := json.Unmarshal([]byte(payload["Body"]), &body)
	if err != nil {
		log.Fatal(err)
	}

	fildsRequeredList := []string{"author", "title", "ups", "num_comments", "created_data"}

	for _, v := range fildsRequeredList {
		if _, ok := body[v]; !ok {
			r := BadRequest{status: 400, message: "field: " + v + " not found"}
			payload["BadRequest"] = utils_api.ToJsonHelper(r)
			utils_api.Logger.Warning(utils_api.ToJsonHelper(payload))
			return false
		}
	}

	return true
}

type BadRequest struct {
	status  int
	message string
}

func ConvertIoReaderToString(data io.Reader) string {
	bodyBytes, err := ioutil.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	return bodyString
}
