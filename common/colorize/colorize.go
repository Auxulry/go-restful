// Package colorize is utilities for colorize script
package colorize

import "fmt"

const (
	Green        = "\u001B[32m"
	Orange       = "\u001B[0;33m"
	Blue         = "\033[34m"
	Red          = "\033[34m"
	DefaultColor = "\033[0m"
)

func MessageColorized(color, m string) string {
	return fmt.Sprintf("%v%v%v", color, m, DefaultColor)
}
