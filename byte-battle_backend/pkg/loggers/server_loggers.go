package loggers

import (
	"fmt"

	"github.com/fatih/color"
)

func StartServerFailure(err error) {
	fmt.Println(color.RedString("ERROR: Could not start the server: %s", err))
}

func StartServerSuccess() {
	fmt.Println("INFO: Server started")
}
