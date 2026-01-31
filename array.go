package main

import "fmt"

func main() {
	var names[3]string
	names[0] = "Bruce"
	names[1] = "Wayne"
	names[2] = "Batman"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int {
		90,
		80,
		96,
		100,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(cap(values))
}