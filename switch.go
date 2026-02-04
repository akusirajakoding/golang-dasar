package main

import "fmt"

func main() {
	name := "Barbara"

	// kode program switch case
	switch name {
	case "Bruce":
		fmt.Println("Hello Master Bruce")
	case "Barbara":
		fmt.Println("Hello Ms Gordon")
	case "Dick":
		fmt.Println("Hello Master Dick")
	case "Jason":
		fmt.Println("Hello Master Todd")
	case "Tim":
		fmt.Println("Hello Master Tim")	
	default:
		fmt.Println("Hi, Introduce yourself")
	}

	// kode program switch case dengan short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("The Name is Too Long")
	case false:
		fmt.Println("The Name is Correct")
	}

	// kode program switch case tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("The Name is Very Long")
	case length > 5:
		fmt.Println("The Name is Long")
	default:
		fmt.Println("The Name is Correct")
	}
}