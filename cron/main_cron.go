package main_cron

import (
	"os"
	"os/signal"
	"time"

	"github.com/robfig/cron/v3"
	cron_services "github.com/sergiodii/cron-go/cron/src/services"
	cron_utils "github.com/sergiodii/cron-go/cron/src/utils"
)

func Main() {

	cron_services.InitStart()

	var cronsRunningList []string

	c := cron.New()
	cron_services.ExecuteCrons(&cronsRunningList, c)

	time.AfterFunc(10*time.Minute, func() {
		cron_services.SyncJobs()
		cron_services.ExecuteCrons(&cronsRunningList, c)
		exChan := make(chan int)
		go func(ch chan int) {
			ch <- 1
		}(exChan)
		<-exChan
	})

	go c.Start()
	cron_utils.Logger.Info("CRON SERVER RUNNIG")
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig

}
