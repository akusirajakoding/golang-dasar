package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var hasilTambah = a + b
	var hasilKurang = b - a
	var hasilKali = a * b
	var hasilBagi = b / a
	var hasilModulus = b % a

	fmt.Println("Hasil Tambah :", hasilTambah)
	fmt.Println("Hasil Kurang :", hasilKurang)
	fmt.Println("Hasil Kali   :", hasilKali)
	fmt.Println("Hasil Bagi   :", hasilBagi)
	fmt.Println("Hasil Modulus:", hasilModulus)
}