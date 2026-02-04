package main

import "fmt"

func main() {
	// kode program for 
	i := 0
	for i < 10 {
		fmt.Println("aku sayang kamu")
		 i++
		fmt.Println(i)
	}

	// for dengan statement
	for i := 0; i < 10; i++ {
		fmt.Println("aku rindu kamu")
	}

	// for dengan array dan slice dengan range
	names := []string{"Bruce", "Wayne", "Batman", "Alfred", "Robin", "Oracle"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
	for _, name := range names {
		fmt.Println("name =", name)
	}
}