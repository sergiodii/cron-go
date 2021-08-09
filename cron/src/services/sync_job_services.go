package cron_services

import (
	"io/ioutil"
	"log"
	"path"
	"reflect"
	"runtime"
	"strings"

	cron_utils "github.com/sergiodii/cron-go/cron/src/utils"
	"github.com/sergiodii/cron-go/domain/entities"
	"github.com/sergiodii/cron-go/domain/use_cases"
	"github.com/sergiodii/cron-go/shared"
)

type jobsSync struct {
	syncInDB   bool
	syncInFile bool
	job        entities.JobsEntity
}
type syncJob struct {
	Name string
	jobsSync
}

func SyncJobs() {
	jobsFromFiles := GetJobsFromFilePath()
	jobsFromDataBase := use_cases.GetJobsUseCase.Execute()

	syncr := make(map[string]jobsSync)
	for _, v := range jobsFromDataBase {
		if s, ok := syncr[v.Name]; ok {
			s.syncInFile = syncr[v.Name].syncInFile
			s.syncInDB = true
			s.job = v
			syncr[v.Name] = s
		} else {
			syncr[v.Name] = jobsSync{syncInDB: true, job: v, syncInFile: false}
		}
	}

	for _, v := range jobsFromFiles {
		if s, ok := syncr[v.Name]; ok {
			s.syncInDB = syncr[v.Name].syncInDB
			s.syncInFile = true
			s.job = v
			syncr[v.Name] = s
		} else {
			syncr[v.Name] = jobsSync{syncInFile: true, job: v, syncInDB: false}
		}
	}

	syncKeys := reflect.ValueOf(syncr).MapKeys()

	var finalList []interface{}

	for _, v := range syncKeys {
		if s, ok := syncr[v.String()]; ok {
			var obj syncJob
			obj.Name = v.String()
			obj.syncInDB = s.syncInDB
			obj.syncInFile = s.syncInFile
			obj.job = s.job
			finalList = append(finalList, obj)
		}
	}

	syncDbChan := make(chan int)
	syncFileChan := make(chan int)

	go func(c chan int) {
		SyncDataIntoFile(finalList)
		c <- 1
	}(syncFileChan)
	go func(c chan int) {
		SyncDataIntoDB(finalList)
		c <- 1
	}(syncDbChan)
	<-syncDbChan
	<-syncFileChan
}

func SyncDataIntoDB(list []interface{}) {
	list = shared.MapArray(shared.FilterArray(list, func(i int, a interface{}, ar []interface{}) bool {
		v, ok := a.(syncJob)
		if !ok {
			return false
		}
		if v.syncInFile && !v.syncInDB {
			return true
		}
		return false
	}), func(i int, a interface{}, ar []interface{}) interface{} {
		v, ok := a.(syncJob)
		if !ok {
			return a
		}
		return v.job
	})
	for _, v := range list {
		f := v.(entities.JobsEntity)
		use_cases.CreateJobUseCase.Execute(f)
	}
}

func SyncDataIntoFile(list []interface{}) {
	list = shared.MapArray(shared.FilterArray(list, func(i int, a interface{}, ar []interface{}) bool {
		v, ok := a.(syncJob)
		if !ok {
			return false
		}
		if !v.syncInFile && v.syncInDB {
			return true
		}
		return false
	}), func(i int, a interface{}, ar []interface{}) interface{} {
		v, ok := a.(syncJob)
		if !ok {
			return a
		}
		return v.job
	})
	__dirname := ""
	if _, file, _, ok := runtime.Caller(0); ok {
		__dirname = path.Dir(file)
	}
	for _, v := range list {
		f := v.(entities.JobsEntity)
		stringJson := []byte(f.ToJson())
		err := ioutil.WriteFile(__dirname+"/../../../"+cron_utils.JobPathString+"/"+strings.ToLower(strings.ReplaceAll(f.Name, "-", ""))+".json", stringJson, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func GetJobsFromFilePath() []entities.JobsEntity {
	fileListFromPath := shared.GetFilesFromPath(cron_utils.JobPathString)
	var jobsFromPath []entities.JobsEntity

	__dirname := ""
	if _, file, _, ok := runtime.Caller(0); ok {
		__dirname = path.Dir(file)
	}

	for _, v := range fileListFromPath {
		jobsFromPath = append(jobsFromPath, ConvertFileToJob(__dirname+"/../../../"+cron_utils.JobPathString+"/"+v.Name()))
	}
	return jobsFromPath
}

func ConvertFileToJob(route string) entities.JobsEntity {
	content, err := ioutil.ReadFile(route)
	if err != nil {
		log.Fatal(err)
	}
	var tempJob entities.JobsEntity
	return *tempJob.FromJson(string(content))
}
