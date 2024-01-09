package outils

import (
	"fmt"
)

const LogEnabled = false

func Printf(format string, a ...any) {
	if LogEnabled {
		fmt.Printf(format, a...)
	}
}

func Println(a ...any) {
	if LogEnabled {
		fmt.Println(a...)
	}
}
