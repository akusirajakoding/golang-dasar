package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	registerUser("dono", func(name string) bool {
		return name == "Anjing"
	})

	filterToxic := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Anjing", filterToxic)
}
