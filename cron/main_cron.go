package main_cron

import (
	"os"
	"os/signal"
	"time"

	"github.com/robfig/cron"
	cron_services "github.com/sergiodii/cron-go/cron/src/services"
	cron_utils "github.com/sergiodii/cron-go/cron/src/utils"
)

func Main() {
	var cronsRunningList []string

	c := cron.New()
	cron_services.ExecuteCrons(&cronsRunningList, c)

	time.AfterFunc(10*time.Minute, func() {
		exChan := make(chan int)
		go func(ch chan int) {
			cron_services.SyncJobs()
			cron_services.ExecuteCrons(&cronsRunningList, c)
			ch <- 1
		}(exChan)
		<-exChan
	})

	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)

	cron_utils.Logger.Fatal("CRON-GO-CRON EXIT:", <-sig)

}

func init() {
	cron_services.SyncJobs()
}
