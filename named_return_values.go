package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Batman"
	middleName = "Bruce"
	lastName = "Wayne"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a, b, c)
}
