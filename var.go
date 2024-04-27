package main

import "fmt"

func main() {
	// var declaration
	var permission bool
	var username string
	fmt.Println(username+"has permission", permission)

	// type convert
	convert := 3.14
	floor := int(convert)
	fmt.Println(convert, "transform type to int is", floor)

	// const and const compute
	const hoursInDay = 24
	const secondsInMinute = 60
	const secondsInDay = secondsInMinute * hoursInDay
	fmt.Println(hoursInDay, secondsInMinute, secondsInDay)

	// format print
	fPnumber := 3.2
	fmt.Printf("fPnumber is %.2f\n", fPnumber)
}
