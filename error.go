package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagian int) (int, error) {
	if pembagian == 0 {
		return 0, errors.New("Hasil Pembagian NOL")
	} else {
		return nilai / pembagian, nil
	}
}

func main() {

	hasil, error := Pembagian(20, 4)
	if error == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", error.Error())
	}

	hasil2, error2 := Pembagian(58, 0)
	if error2 == nil {
		fmt.Println("Hasil", hasil2)
	} else {
		fmt.Println("Error", error2.Error())
	}

}
