package ui

import "fmt"

func ShowMessage(message string) {
	fmt.Println(message)
}

func ShowFormattedMessage(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
