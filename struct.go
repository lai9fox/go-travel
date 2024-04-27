package main

import "fmt"

type User struct {
	name       string
	permission string
}

type UserGroup struct {
	name  string
	users []User
}

type rect struct {
	height, width int
}

func (r *rect) area() int {
	return r.height * r.width
}

func main() {
	admin := User{"admin", "admin"}
	guest := User{"guest", "guest"}
	anonymous := User{}
	fmt.Println(admin, guest, anonymous)

	// anonymous struct
	car := struct {
		name  string
		price int
	}{
		"geely",
		169000,
	}
	fmt.Println("anonymous struct", car)

	//	method on struct
	r := rect{height: 10, width: 5}
	fmt.Println("rect", r.area())
}
