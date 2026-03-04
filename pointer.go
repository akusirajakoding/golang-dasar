package main 

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1 // address2 adalah copy dari address1, jadi address1 dan address2 memiliki data yang sama

	// address2.City = "Bandung"

	// fmt.Println(address1) // address1 tidak berubah karena address2 adalah copy dari address1
	// fmt.Println(address2) // address2 berubah karena kita mengubah city di address2

	// jika kita ingin membuat address2 yang merubah city tapi tidak merubah city di address1, maka kita bisa menggunakan pointer
	// pointer adalah sebuah variable yang menyimpan alamat memory dari variable lain
	// dengan menggunakan pointer, kita bisa merubah data di address1 melalui address2
	// untuk membuat pointer, kita bisa menggunakan operator & untuk mendapatkan alamat memory dari variable
	// dan untuk mengakses data dari pointer, kita bisa menggunakan operator * untuk mendapatkan data dari alamat memory yang disimpan di pointer

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // address2 adalah pointer yang menyimpan alamat memory dari address1, jadi address2 dan address1 memiliki data yang sama

	address2.City = "Bandung"

	fmt.Println(address1) // address1 berubah karena address2 adalah pointer yang menyimpan alamat memory dari address1
	fmt.Println(address2) 
}