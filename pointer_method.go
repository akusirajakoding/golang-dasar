package main

import "fmt"

type Man struct {
	Name string	
}

// untuk method direkomendasikan menggunakan pointer agar bisa merubah data structnya

func (man *Man) Married() { // cukup tambahkan pointer agar bisa merubah data structnya
	man.Name = "Mr. " + man.Name
}

func main() {
	joko := Man{Name: "Joko"}
	joko.Married()

	fmt.Println(joko.Name)
}