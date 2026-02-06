package main

import "fmt"

func getFullName()(string, string) {
	return "Bruce", "Wayne"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}