package main

import "fmt"

func main() {
	const NamaLengkap = "Bruce Wayne"
	fmt.Println("Nama Lengkap :", NamaLengkap)

	const (
		FirstName = "Bruce"
		LastName  = "Wayne"
	)
	fmt.Println(FirstName)
	fmt.Println(LastName)
}