package utils

import (
	"os"
	"strings"
)

type AnsiColor string

const (
	BLACK        AnsiColor = "\033[0;30m"
	RED          AnsiColor = "\033[0;31m"
	GREEN        AnsiColor = "\033[0;32m"
	BROWN        AnsiColor = "\033[0;33m"
	BLUE         AnsiColor = "\033[0;34m"
	PURPLE       AnsiColor = "\033[0;35m"
	CYAN         AnsiColor = "\033[0;36m"
	LIGHT_GRAY   AnsiColor = "\033[0;37m"
	DARK_GRAY    AnsiColor = "\033[1;30m"
	LIGHT_RED    AnsiColor = "\033[1;31m"
	LIGHT_GREEN  AnsiColor = "\033[1;32m"
	YELLOW       AnsiColor = "\033[1;33m"
	LIGHT_BLUE   AnsiColor = "\033[1;34m"
	LIGHT_PURPLE AnsiColor = "\033[1;35m"
	LIGHT_CYAN   AnsiColor = "\033[1;36m"
	LIGHT_WHITE  AnsiColor = "\033[1;37m"
	BOLD         AnsiColor = "\033[1m"
	FAINT        AnsiColor = "\033[2m"
	ITALIC       AnsiColor = "\033[3m"
	UNDERLINE    AnsiColor = "\033[4m"
	BLINK        AnsiColor = "\033[5m"
	NEGATIVE     AnsiColor = "\033[7m"
	CROSSED      AnsiColor = "\033[9m"
	RESET        AnsiColor = "\033[0m"
)

func Colorize(input string, targets []string, color AnsiColor) string {
	for _, target := range targets {
		colorized := string(color) + `"` + target + `"` + string(RESET)
		input = strings.Replace(input, `"`+target+`"`, colorized, 1)
	}
	return input
}

func GetEnvVar(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
