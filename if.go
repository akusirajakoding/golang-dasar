package main

import "fmt"

func main() {
	name := "Barbara"

	if name == "Bruce" {
		fmt.Println("Hello Master Bruce")
	} else if name == "Dick" {
		fmt.Println("Hello Master Dick")
	} else if name == "Barbara" {
		fmt.Println("Hello Ms Gordon")
	} else if name == "Jason" {
		fmt.Println("Hello Master Todd")
	} else if name == "Tim" {
		fmt.Println("Hello Master Tim")
	} else {
		fmt.Println("Hi, Introduce yourself")
	}

	if length := len(name); length > 5 {
		fmt.Println("The Name is Too Long")
	} else {
		fmt.Println("The Name is Correct")
	}
}