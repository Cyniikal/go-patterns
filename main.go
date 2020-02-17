package main

import (
	"fmt"
	"go-patterns/decorator"
)

func main() {
	// Create a StringRunner, add middleware, and run
	sc := decorator.StringRunner{Middleware: make([]decorator.StringMiddleware, 0)}
	sc.AddMiddleware(decorator.ReverseString, decorator.ToLower)
	fmt.Println(sc.Run("Hello, WORLD!"))
}
