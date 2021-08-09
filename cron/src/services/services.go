package cron_services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/novalagung/golpal"
	"github.com/robfig/cron"
	"github.com/sergiodii/cron-go/domain/use_cases"
	"github.com/sergiodii/cron-go/shared"
)

func ExecuteCrons(atualCronList *[]string, cron *cron.Cron) {
	jobList := use_cases.GetJobsUseCase.Execute()
	var interList []interface{}
	for _, jobName := range *atualCronList {
		interList = append(interList, jobName)
	}
	for _, job := range jobList {
		runnig := shared.FindArray(interList, func(ind int, item interface{}, ar []interface{}) bool {
			j := item.(string)
			fmt.Println(j, job.Name)
			if j == job.Name {
				return true
			}
			return false
		})
		if runnig == nil {
			cron.AddFunc(job.Cron, func() {
				cmdString := GetRepoJobs("https://raw.githubusercontent.com/sergiodii/cron-go-jobs-function/master/main.go")
				handleFunction := job.Function + "()"
				if len(job.Parans) >= 1 {
					handleFunction = job.Function + "(" + strings.Join(job.Parans, ", ") + ")"
				}
				cmdString = strings.ReplaceAll(cmdString, "//##localToExecution##", handleFunction)
				output, err := golpal.New().ExecuteRaw(cmdString)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("result", "=>", output)
			})
			*atualCronList = append(*atualCronList, job.Name)
		}
	}
}

func GetRepoJobs(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		return bodyString
	}
	return ""
}
