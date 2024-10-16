package sitecore

import (
	"fmt"
)

type LogLevel string

const (
	Editing  LogLevel = "Editing"
	Renderer LogLevel = "Renderer"
	All      LogLevel = "*"
)

var (
	loggerMap = map[string]LogLevel{
		"Editin":   Editing,
		"Renderer": Renderer,
		"*":        All,
	}
)

func getLogger(logger string) LogLevel {
	return loggerMap[logger]
}

func EditingLog(message string, a ...any) {
	logMessage(Editing, message, a...)
}

func RendererLog(message string, a ...any) {
	logMessage(Renderer, message, a...)
}

func logMessage(logger LogLevel, message string, a ...any) {

	enabledLogger := getLogger(GetEnvVar("SITECORE_DEBUG_LOGGER"))
	if enabledLogger == logger || enabledLogger == All {
		fmt.Printf(message, a...)
	}
}
