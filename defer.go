package main

import "fmt"

func logging() {
	fmt.Println("End App")
}

func runApplication() {
	defer logging()
	fmt.Println("hello world")
}

func main() {
	runApplication()
}
