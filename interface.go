package main

import "fmt"

type animal interface {
	takeFood()
	move(x int, y int)
}

type bird struct {
	name   string
	color  string
	recipe string
}

func (b bird) takeFood() {
	fmt.Printf("%s is %s, take in %s", b.name, b.color, b.recipe)
}

func (b bird) move(x int, y int) {
	fmt.Printf("%s is flying to %d,%d.", b.name, x, y)
}

type dog struct {
	name string
}

func (d dog) takeFood() {
	fmt.Printf("%s take anything", d.name)
}

func (d dog) move(x int, y int) {
	fmt.Printf("%s run to %d,%d.", d.name, x, y)
}

func main() {
	var a animal

	a = dog{"keji"}
	a.move(10, 20)

	a = bird{
		name:   "beke",
		color:  "green",
		recipe: "bugs",
	}
	a.takeFood()
	a.move(20, 20)
}
