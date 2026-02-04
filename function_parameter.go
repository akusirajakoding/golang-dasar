package main 

import "fmt"

func sayHelloTo(firstName string, lastName string, age int) {
	fmt.Println("Hello", firstName, lastName, "your age", age)
}

func main() {
	sayHelloTo("Bruce", "Wayne", 30)
}