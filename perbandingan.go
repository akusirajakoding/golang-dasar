package main

import "fmt"

func main() {
	var bahasa1 = "Golang"
	var bahasa2 = "Python"
	var isSame = bahasa1 == bahasa2
	var isDifferent = bahasa1 != bahasa2

	fmt.Println("Bahasa 1 :", bahasa1)
	fmt.Println("Bahasa 2 :", bahasa2)
	fmt.Println("Apakah bahasa1 dan bahasa2 berbeda?", isDifferent)
	fmt.Println("Apakah bahasa1 dan bahasa2 sama?", isSame)
}