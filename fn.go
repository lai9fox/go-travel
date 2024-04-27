package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func getPoint() (int, int) {
	return 3, 4
}

func main() {
	// add fn
	x := 1
	y := 2
	fmt.Println("x plus y is", add(x, y))

	// multi return value
	xP, yP := getPoint()
	fmt.Println(xP, yP)
}
