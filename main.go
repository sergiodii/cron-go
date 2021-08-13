package main

import (
	"os"

	main_api "github.com/sergiodii/cron-go/api"
	main_cron "github.com/sergiodii/cron-go/cron"
	"github.com/sergiodii/cron-go/shared"
)

func main() {

	args := shared.GetArgs(os.Args)

	if _, ok := args["api"]; ok {
		main_api.Main()
		return
	}

	if _, ok := args["cron"]; ok {
		main_cron.Main()
	}
}
