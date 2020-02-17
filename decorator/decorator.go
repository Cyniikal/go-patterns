package decorator

import "strings"

type StringRunner struct {
	Middleware []StringMiddleware
}

func (sc *StringRunner) AddMiddleware(m ...StringMiddleware) *StringRunner {
	for _, middleware := range m {
		sc.Middleware = append(sc.Middleware, middleware)
	}
	return sc
}
func (sc *StringRunner) Run(s string) string {
	runFunc := func(s string) string { return s }
	for _, middleware := range sc.Middleware {
		runFunc = middleware(runFunc)
	}
	return runFunc(s)
}

type StringManipulator func(string) string
type StringMiddleware func(StringManipulator) StringManipulator

func ReverseString(sm StringManipulator) StringManipulator {
	return func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}
}

func ToLower(sm StringManipulator) StringManipulator {
	return func(s string) string {
		return sm(strings.ToLower(s))
	}
}