package cron_services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/novalagung/golpal"
	"github.com/robfig/cron/v3"
	cron_config "github.com/sergiodii/cron-go/cron/src/config"
	cron_utils "github.com/sergiodii/cron-go/cron/src/utils"
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
    
		if runnig == nil || runnig == false {
			cron_utils.Logger.Info("CRON-ADD JSON: ", job.ToJson())
			cmdString := GetRepoJobs(cron_config.JobsHandleGithub)
			handleFunction := job.Function + "()"
			if len(job.Parans) >= 1 {
				handleFunction = job.Function + "(" + strings.Join(job.Parans, ", ") + ")"
			}
			cmdString = strings.ReplaceAll(cmdString, "//##localToExecution##", handleFunction)

			output, err := golpal.New().ExecuteRaw(cmdString)
			if err != nil {
				cron_utils.Logger.Error(err)
			}
			cron_utils.Logger.Info("CRON-EXECUTION RESULT =>", output)

			cron.AddFunc(job.Cron, func() {
				cron_utils.Logger.Info("CRON-ADD EXECUTION: ", job.ToJson())
				output, err := golpal.New().ExecuteRaw(cmdString)
				if err != nil {
					cron_utils.Logger.Error(err)
				}
				cron_utils.Logger.Info("CRON-EXECUTION RESULT =>", output)
			})
			*atualCronList = append(*atualCronList, job.Name)
		}
	}
}

func GetRepoJobs(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		cron_utils.Logger.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			cron_utils.Logger.Error(err)
		}
		bodyString := string(bodyBytes)
		return bodyString
	}
	return ""
}

func InitStart() {
	//Migrate Job table if not exits
	use_cases.MigrateTableJobUseCase.Execute()
	SyncJobs()
}
