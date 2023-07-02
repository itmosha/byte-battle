package loggers

import (
	"fmt"

	"github.com/fatih/color"
)

func LoadEnvFailure() {
	fmt.Println(color.RedString("ERROR: Could not load .env file"))
}

func VariableNotFound(variable string) {
	fmt.Println(color.RedString("ERROR: Variable %s was not found in .env file", variable))
}
