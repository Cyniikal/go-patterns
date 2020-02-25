package main

import (
	"fmt"
	"go-patterns/decorator"
	"go-patterns/facade"
	"go-patterns/factory"
	"go-patterns/strategy"
	"go-patterns/wrapper"
)

func main() {
	// Decorator
	// Create a StringRunner, add middleware, and run
	// Decorators let us add additional functionality that runs before the base function.
	// Specifically, they let us modify the input and perform actions before and after the base function is executed.
	sc := decorator.StringRunner{Middleware: make([]decorator.StringMiddleware, 0)}
	sc.AddMiddleware(decorator.ReverseString, decorator.ToLower)
	fmt.Println("Decorator", sc.Run("Hello, WORLD!"))
	fmt.Println()

	// Wrapper
	// Provide multiple different functions to make calling a complex function easier for general users of the API
	fmt.Println("Wrapper:", wrapper.ModeOne("WRAPPER_TEST"))
	fmt.Println("Wrapper:", wrapper.ModeTwo("WRAPPER_TEST"))
	fmt.Println()

	// Strategy
	// Create a strategy and pass it different tactics to get different operations at runtime
	strat := strategy.Tactic{Strategy: strategy.CapsTactic{}}
	fmt.Println("Strategy:", strat.Apply("Hello, World!"))
	strat.Strategy = strategy.LowerTactic{}
	fmt.Println("Strategy:", strat.Apply("Hello, World!"))
	fmt.Println()

	// Factory
	// Returns a given type of object without use having to create them directly
	// Factories also allow us to hide the complexities of instantiation and initialization from users of our API
	vanillaIceCream := factory.NewIceCream(factory.Vanilla)
	chocolateIceCream := factory.NewIceCream(factory.Chocolate)
	fmt.Println("Factory:", vanillaIceCream)
	fmt.Println("Factory:", chocolateIceCream)
	fmt.Println()

	// Facade
	// Mask complex interactions of types with an easy-to-use API
	// Similar to Factory in that it can be used to help ease initialization, but the intent is slightly different
	complexThing := facade.CreateD(5)
	fmt.Println("Facade:", facade.GetBaseVal(complexThing))
	_ = facade.UpdateBaseVal(complexThing)
	fmt.Println("Facade:", facade.GetBaseVal(complexThing))
}
