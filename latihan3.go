package main

import "fmt"

func cetakNilai(nilai int) {
	if nilai >= 75 {
		fmt.Println("LULUS")
	} else if nilai < 75 {
		fmt.Println("TIDAK LULUS")
	}
}

func tambah(a int, b int) int {
	return a + b
}

func cariAngkaLima (data []int) int {
	for _, value := range data {
		if value == 5 {
			return value
		}
	}
	return -1
}

func main() {
	cetakNilai(80)
	cetakNilai(60)
	fmt.Println(" ")
	
	fmt.Println("Hasil function tambah:", tambah(10, 20))
	result := tambah(15, 25)
	fmt.Println("Hasil function tambah:", result)
	fmt.Println(" ")

	nilai := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	hasil := cariAngkaLima(nilai)
	fmt.Println("Hasil function cariAngkaLima:", hasil)


}