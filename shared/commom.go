package shared

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// package switch expression {
// case condition:

// }
func GetArgs(a []string) map[string]interface{} {
	final := make(map[string]interface{})
	for i, v := range a {
		if i >= 1 {
			validParam := regexp.MustCompile(`^[-][-]?`)
			regexKey := regexp.MustCompile(`^[-][-]?`)
			if validParam.MatchString(v) {
				final[regexKey.ReplaceAllString(v, "")] = true
			} else if !validParam.MatchString(v) && validParam.MatchString(a[i-1]) {
				final[regexKey.ReplaceAllString(a[i-1], "")] = v
			}
		}
	}
	return final
}

func GetEnvFileName() string {
	args := GetArgs(os.Args)
	if v, ok := args["env"]; ok {
		return fmt.Sprintf("%s", v)
	}
	return ".env"
}

func GetEnvOrFail(env string) string {
	e := os.Getenv(env)
	if len(e) <= 0 {
		log.Fatal("\033[31mERROR - GET_ENV_OR_FAIL: \033[33m" + env + "\033[31m Not Found\033[0m")
	}
	return e
}

func printFatal(list ...interface{}) {
	logger := NewLogger("")
	defer logger.CloseFile()
	logger.Fatal(list)
}

// GetFilesFromPath - Get the files infos, from the root folder ex: cron/src...
func GetFilesFromPath(direction string) []fs.FileInfo {
	__dirname := ""
	if _, file, _, ok := runtime.Caller(0); ok {
		__dirname = path.Dir(file)
	} else {
		printFatal("error locating logger directory")
		return nil
	}
	files, err := ioutil.ReadDir(__dirname + "/../" + direction)
	if err != nil {
		printFatal(err)
	}
	return files
}

func FilterArray(array []interface{}, callback func(int, interface{}, []interface{}) bool) []interface{} {
	var filtered []interface{}
	for i, v := range array {
		if callback(i, v, array) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func MapArray(array []interface{}, callback func(int, interface{}, []interface{}) interface{}) []interface{} {
	var maped []interface{}
	for i, v := range array {
		maped = append(maped, callback(i, v, array))
	}
	return maped
}

func FindArray(array []interface{}, callback func(int, interface{}, []interface{}) bool) interface{} {
	var item interface{}
	returnNull := true
	for i, v := range array {
		if callback(i, v, array) {
			returnNull = false
			item = v
			break
		}
	}
	if returnNull {
		return nil
	}
	return item
}

type TimesTamp struct {
	integer int64
}

func (t *TimesTamp) UnmarshalJSON(data []byte) error {
	in, err := strconv.ParseInt(strings.ReplaceAll(string(data[:]), ".0", ""), 10, 64)
	if err != nil {
		return err
	}
	t.integer = in
	// t.text = strings.ReplaceAll(string(data[:]), ".0", "")
	return nil
}

func (t *TimesTamp) ToInt() int64 {
	return t.integer
}
