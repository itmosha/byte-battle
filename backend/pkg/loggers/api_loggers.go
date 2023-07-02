package loggers

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func ApiRequestSuccess(method string, path string, status int) {
	fmt.Printf("API REQUEST: %s /api%s [%d %s]\n", method, path, status, http.StatusText(status))
}

func ApiRequestFailure(method string, path string, status int) {
	fmt.Println(color.RedString("API REQUEST: %s /api/%s [%d %s]", method, path, status, http.StatusText(status)))
}
