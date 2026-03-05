package main

import "fmt"

//type Address struct {
//	City, Province, Country string
//}
//
//func ChangeCountryToIndonesia(address Address) {
//	address.Country = "Indonesia"
//}
//
//func main() {
//	address := Address{}
//	ChangeCountryToIndonesia(address)
//
//	fmt.Println(address)
//
//	// dengan menulis kode seperti di atas tidak akan mengubah address karena tidak ada pointer
//
//}

// dengan menulis kode seperti di bawah ini akan mengubah address karena ada pointer

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) { // ada pointer
	address.Country = "Indonesia"
}

func main() {
	address := Address{}
	ChangeCountryToIndonesia(&address) // ada pointer

	fmt.Println(address)
}
