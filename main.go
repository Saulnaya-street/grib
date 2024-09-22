package main

import "fmt"

type Animal interface {
	Sound() string
	Move() string
}
type Dog struct{}

func (d Dog) Sound() string {
	return "Гав"
}

func (d Dog) Move() string {
	return "Рычит"
}

type Cat struct{}

func (c Cat) Sound() string {
	return "Мау"
}

func (c Cat) Move() string {
	return "Царапает"
}
func zvAnimal(a Animal) {
	fmt.Printf("Издает звук: %s делает действие %s \n", a.Sound(), a.Move())

}

func main() {
	dog := Dog{}
	cat := Cat{}

	zvAnimal(dog)
	zvAnimal(cat)
}
