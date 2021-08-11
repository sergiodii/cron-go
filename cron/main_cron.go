package main_cron

import (
	"os"
	"os/signal"
	"time"

	"github.com/robfig/cron"
	cron_services "github.com/sergiodii/cron-go/cron/src/services"
)

func Main() {

	cron_services.SyncJobs()

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
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
	// cron_utils.Logger.Fatal("CRON-GO-CRON EXIT:", <-sig)

}
