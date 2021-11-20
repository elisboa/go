package main

import "fmt"

func FullName(firstName string, lastName string) string {

	fullName := firstName + " " + lastName
	return fullName
}

func main() {

	fmt.Println("Hello, world!")

	fullName := FullName("Eduardo", "Lisboa")
	fmt.Println("My full name is", fullName)
}
