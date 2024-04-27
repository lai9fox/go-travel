package main

import (
	"errors"
	"fmt"
)

type MyCustomError interface {
	Error() string
}

type divide struct {
}

func (d divide) Error() string {
	return "divide error"
}

func main() {
	// custom error
	d := divide{}
	fmt.Println(d.Error())

	// pkg error
	fmt.Println(errors.New("official error package"))
}
