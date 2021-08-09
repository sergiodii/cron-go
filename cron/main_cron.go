package main_cron

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/robfig/cron"
	cron_config "github.com/sergiodii/cron-go/cron/src/config"
	cron_services "github.com/sergiodii/cron-go/cron/src/services"
)

func Main() {
	var cronsRunningList []string
	cronsRunningList = append(cronsRunningList, "call-reddit")
	c := cron.New()
	cron_services.ExecuteCrons(&cronsRunningList, c)
	fmt.Println(cronsRunningList)
	c.AddFunc("* * * * * *", RunEverySecond)
	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)

	fmt.Println("sergio", <-sig)

}

func RunEverySecond() {
	fmt.Printf("%v\n", time.Now())
}
func RunEverySecond2() {
	fmt.Printf("%s: %v\n", "SERGIO", time.Now())
}
func init() {
	cron_config.StartConfig()
}
