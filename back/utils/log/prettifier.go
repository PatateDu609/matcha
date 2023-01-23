package log

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

var colors map[string]string = map[string]string{
	"file fragment":          "\033[38;5;83m",
	"function fragment":      "\033[38;5;41m",
	"numbered init fragment": "\033[38;5;47m",
	"package fragment":       "\033[38;5;70m",
	"function dots":          "\033[38;5;111m",
	"file slashes":           "\033[38;5;10m",
	"colon":                  "\033[38;5;48m",
	"line number":            "\033[1;38;5;45m",
	"reset":                  "\033[0m",
}

func Prettyfier(frame *runtime.Frame) (function string, file string) {
	function = frame.Function
	file = frame.File

	if strings.HasPrefix(function, "github.com/PatateDu609/matcha/") {
		function = strings.TrimPrefix(function, "github.com/PatateDu609/matcha/")
	}

	if strings.HasPrefix(file, "/app/") {
		file = strings.TrimPrefix(file, "/app/")
	} else if strings.HasPrefix(file, "/Users/ghaliboucetta/Workspace/matcha/") {
		file = strings.TrimPrefix(file, "/Users/ghaliboucetta/Workspace/matcha/")
	}

	functionFragments := strings.Split(function, ".")
	lastFragment := functionFragments[len(functionFragments)-1]
	numberedInit := false
	if _, err := strconv.ParseInt(lastFragment, 10, 64); err == nil {
		numberedInit = true
	}
	for i, fragment := range functionFragments {
		color := colors["package fragment"]

		if numberedInit && i == len(functionFragments)-1 {
			color = colors["numbered init fragment"]
		} else if (numberedInit && i == len(functionFragments)-2) || (!numberedInit && i == len(functionFragments)-1) {
			color = colors["function fragment"]
		}

		functionFragments[i] = fmt.Sprintf("%s%s%s", color, fragment, colors["reset"])
	}
	function = strings.Join(functionFragments, fmt.Sprintf("%s.%s", colors["function dots"], colors["reset"]))

	fileFragments := strings.Split(file, "/")
	for i, fragment := range fileFragments {
		fileFragments[i] = fmt.Sprintf("%s%s%s", colors["file fragment"], fragment, colors["reset"])
	}
	file = strings.Join(fileFragments, fmt.Sprintf("%s/%s", colors["file slashes"], colors["reset"]))

	file = fmt.Sprintf("%s%s:%s%d%s", file, colors["colon"], colors["line number"], frame.Line, colors["reset"])
	return
}
