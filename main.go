package main

import (
	"fmt"
	"os"

	main_api "github.com/sergiodii/cron-go/api"
	main_cron "github.com/sergiodii/cron-go/cron"
	"github.com/sergiodii/cron-go/shared"
)

func task() {
	fmt.Println("I am running task.")
}

func task2() {
	fmt.Println("I am running task.2")
}

func main() {

	// channel1 := make(chan int)
	// channel2 := make(chan int)

	// go func(ch chan int) {
	// 	time.Duration.Seconds(3)
	// 	ch <- 1
	// }(channel1)

	// time.Duration.Seconds(1)

	// go func(ch chan int) {
	// 	time.Duration.Seconds(5)
	// 	ch <- 2
	// }(channel2)

	// fmt.Println(<-channel2)
	// fmt.Println(<-channel1)

	// files := cron_services.GetFilesFromPath("cron")
	// for _, f := range files {
	// 	fmt.Println("Name: ", f.Name())
	// 	fmt.Println("Is Dir?: ", f.IsDir())
	// }

	args := shared.GetArgs(os.Args)

	if _, ok := args["api"]; ok {
		main_api.Main()
		return
	}

	if _, ok := args["cron"]; ok {
		main_cron.Main()
	}

}
