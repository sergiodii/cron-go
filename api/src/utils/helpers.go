package utils_api

import (
	"encoding/json"
	"os"

	cli_commands "github.com/sergiodii/cron-go/shared/cli/commands"
)

func ToJsonHelper(v interface{}) string {
	b, e := json.Marshal(v)
	if e != nil {
		return `{"error":"error"}`
	}
	return string(b)
}

func GetServerPortHelper() string {
	port := cli_commands.EnvFlags.Port()
	portEnv := os.Getenv("PORT")
	if len(portEnv) >= 1 && *port == cli_commands.DefaultPort {
		return ":" + portEnv
	}
	return ":" + *port
}
