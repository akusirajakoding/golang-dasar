package helper

var version = "1.0.0" // hanya bisa diakses didalam package helper saja, tidak bisa diakses dari package lain
var Application = "Golang Dasar" // bisa diakses dari package lain, karena diawali dengan huruf besar

func sayGoodbye(name string) string { // jika membuat function dengan huruf kecil, maka function tersebut hanya bisa diakses didalam package yang sama
	return "Goodbye " + name
}

func SayHello(name string) string { // jika membuat function dengan huruf besar, maka function tersebut bisa diakses dari package lain
	return "Hello " + name
}
