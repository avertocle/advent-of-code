package clr

import "fmt"

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func Str(s string, color string) string {
	return fmt.Sprintf("%v%v%v", color, s, Reset)
}

func Int(i int, color string) string {
	s := fmt.Sprintf("%v", i)
	return Str(s, color)
}

func Gen(i interface{}, color string) string {
	s := fmt.Sprintf("%v", i)
	return Str(s, color)
}
