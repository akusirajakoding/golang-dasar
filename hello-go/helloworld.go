package main

import "fmt"

fn func main() {
	names := []string{"anto", "dewi", "nanda", "junior"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}
