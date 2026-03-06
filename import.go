package main

import (
	"fmt"
	"golang-dasar/helper"
)

func main() {
	result := helper.SayHello("Dono")
	fmt.Println(result)

	fmt.Println(helper.Application) // bisa mengakses variable Application yang ada di package helper, karena diawali dengan huruf besar
	fmt.Println(helper.version) // tidak bisa mengakses variable version yang ada di package helper, karena diawali dengan huruf kecil
	fmt.Println(helper.sayGoodbye("Dono")) // tidak bisa mengakses function sayGoodbye yang ada di package helper, karena diawali dengan huruf kecil

}
