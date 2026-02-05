package main 

import "fmt"

func cekLulus(nilai int) {
	if nilai >= 75 {
		fmt.Println("LULUS")
	} else if nilai < 75 {
		fmt.Println("TIDAK LULUS")
	}
}

func main() {
	cekLulus(80)
	cekLulus(60)
	cekLulus(75)
	cekLulus(74)
}