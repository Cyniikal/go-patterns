package main

import (
	"fmt"
	"go-patterns/decorator"
	"go-patterns/strategy"
)

func main() {
	// Decorator
	// Create a StringRunner, add middleware, and run
	sc := decorator.StringRunner{Middleware: make([]decorator.StringMiddleware, 0)}
	sc.AddMiddleware(decorator.ReverseString, decorator.ToLower)
	fmt.Println("Decorator", sc.Run("Hello, WORLD!"))

	// Strategy
	// Create a strategy and pass it different tactics to get different operations at runtime
	strat := strategy.Tactic{Strategy: strategy.CapsTactic{}}
	fmt.Println("Strategy:", strat.Apply("Hello, World!"))
	strat.Strategy = strategy.LowerTactic{}
	fmt.Println("Strategy:", strat.Apply("Hello, World!"))
}
