package main

import "fmt"

func test(orang1 string, orang2 string, orang3 string, orang4 string) (string, string, string, string) {
	return orang1, orang2, orang3, orang4
}

func main() {
	orang1, orang2, orang3, orang4 := test("ando", "nano", "jono", "dono")
	fmt.Println(orang1, orang2, orang3, orang4)

	fmt.Println(test("dono", "tono", "toni", "doni"))

	orang5, orang6, orang7, orang8 := test("ino", "oni", "nio", "noi")
	fmt.Println(orang5, orang6, orang7, orang8)
}
