package main

import "fmt"

func main() {
	name := "Bruce"
	fmt.Println("Nama Saya :", name)

	name = "Bruce Wayne"
	fmt.Println("Nama Saya :", name)

	var (
		firstName = "Bruce"
		lastName  = "Wayne"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}