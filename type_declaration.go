package main

import "fmt"

func main() {
	type NoKTP string

	var KTP NoKTP = "3275023405670001"
	fmt.Println("No KTP :", KTP)

	var contohString string = "324234141511550001"
	var contohKTP NoKTP = NoKTP(contohString)
	fmt.Println("Contoh KTP :", contohKTP)
}