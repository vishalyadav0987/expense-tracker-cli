package cli

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	successColor = color.New(color.FgGreen).SprintFunc()
	errorColor   = color.New(color.FgRed).SprintFunc()
	warnColor    = color.New(color.FgYellow).SprintFunc()
	infoColor    = color.New(color.FgCyan).SprintFunc()
)

func PrintSuccess(msg string) {
	fmt.Println(successColor("✔ " + msg))
}

func PrintError(msg error) {
	fmt.Println(errorColor("✖ ", msg))
}

func PrintWarning(msg string) {
	fmt.Println(warnColor("⚠ " + msg))
}

func PrintInfo(msg string) {
	fmt.Println(infoColor("ℹ " + msg))
}
