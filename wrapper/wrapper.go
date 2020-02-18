package wrapper

import "fmt"

func baseFunc(s string, mode int) string {
	var r string
	switch mode {
	case 1:
		r = "Mode 1"
	case 2:
		r = "Mode 2"
	default:
		r = "Unsupported Mode"
	}
	return fmt.Sprintf("baseFunc called with %s and given string {%s}", r, s)
}

func ModeOne(s string) string {
	return baseFunc(s, 1)
}

func ModeTwo(s string) string {
	return baseFunc(s, 2)
}