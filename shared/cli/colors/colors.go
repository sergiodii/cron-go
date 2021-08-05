package cli_colors

import (
	"fmt"
)

type CliColor string

var Reset CliColor = "\033[0m"
var Red CliColor = "\033[31m"
var Green CliColor = "\033[32m"
var Yellow CliColor = "\033[33m"
var Blue CliColor = "\033[34m"
var Purple CliColor = "\033[35m"
var Cyan CliColor = "\033[36m"
var White CliColor = "\033[37m"

func PrintColor(color CliColor, text string) {
	fmt.Println(string(color), text, string(Reset))
}
