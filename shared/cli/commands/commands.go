package cli_commands

import (
	"fmt"
	"os"

	"github.com/sergiodii/cron-go/shared"
)

var DefaultPort string = "8080"

type flags struct {
	port *string
}

func (f *flags) Init() {
	args := shared.GetArgs(os.Args)
	f.port = &DefaultPort
	if v, ok := args["port"]; ok {
		p := fmt.Sprintf("%v", v)
		f.port = &p
	}
}

func (f *flags) Port() *string {
	if f.port == nil {
		f.Init()
	}
	return f.port
}

var EnvFlags flags
