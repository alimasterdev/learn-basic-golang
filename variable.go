package main

import (
	"fmt"
)

func main() {
	var name string
	name = "Muhammad Ali Imron"
	fmt.Println("Fullname = ", name)

	var age int8
	age = 25
	fmt.Println("Age = ", age)

	country := 55
	fmt.Println("Country = ", country)

	var (
		firstName = "Muhammad Ali "
		lastName  = "Imron"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println("Fullname = ", firstName+lastName)
}
