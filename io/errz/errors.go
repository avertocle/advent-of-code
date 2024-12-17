package errz

import (
	"fmt"
	"log"
)

func SoftAssert(condition bool, format string, a ...any) {
	if !condition {
		fmt.Printf(format, a...)
		fmt.Println()
	}
}

func HardAssert(condition bool, format string, a ...any) {
	if !condition {
		log.Fatalf(format, a...)
	}
}
