package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

/*
struct adalah template data atau prototype data
struct tidak dapat langsung digunakan
kita bisa membuat data/object dari struct yg sudah kita buat
*/

// struct method
/*
struct adalah tipe data seperti tipe data lainnya, bisa digunakan sebagai parameter untuk function
namun jika kita ingin menambah method ke dalam struct, sehingga seakan akan sebuah struct memiliki function
maka method adalah function
*/

// method function
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var bruce Customer
	fmt.Println(bruce)

	bruce.Name = "Bruce Wayne"
	bruce.Address = "Gotham"
	bruce.Age = 30
	fmt.Println(bruce)
	fmt.Println(bruce.Name)
	fmt.Println(bruce.Address)
	fmt.Println(bruce.Age)

	alfred := Customer{
		Name:    "Alfred Pennyworth",
		Address: "Gotham",
		Age:     60,
	}
	fmt.Println(alfred)

	robin := Customer{"Tim Drake", "Gotham", 20}
	fmt.Println(robin)

	bruce.sayHello("Barbara")
	alfred.sayHello("Barbara")
	robin.sayHello("Barbara")

}
