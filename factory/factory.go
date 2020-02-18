package factory

import "fmt"

type Food interface {
	Eat() bool
}

type Flavor int
func (f Flavor) Eat() bool {
	return true
}

const (
	Vanilla Flavor = iota + 1
	Chocolate
)

func NewIceCream(flavor Flavor) (created Food) {
	switch flavor {
	case Vanilla:
		fmt.Println("Factory: Doing complex stuff to create vanilla ice cream")
		created = Vanilla
	case Chocolate:
		fmt.Println("Factory: Doing complex stuff to create chocolate ice cream")
		created = Chocolate
	}
	return
}