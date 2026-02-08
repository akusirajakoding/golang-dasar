package main

import "fmt"

func getGoodBye(name string) string {
	return "Goodbye " + name
}

func main() {
	goodbye := getGoodBye // jika ingin menjadikan function sebagai value jangan gunakan () di akhir
	// nama function nya, kalo gitu jadinya manggil si function bukan mau ngejadiin function sebagai value
	fmt.Println(goodbye("Master Bruce"))
}
