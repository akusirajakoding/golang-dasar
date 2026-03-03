package main

import "fmt"

func NewMap (name string) map[string]string {
	if name == "" {	
		return nil
	} else {
		return map[string]string{
			"name": name,
	}
}
}

func main() {
	data1 := NewMap("")

	if data1 == nil {
		fmt.Println("Data masih kosong")
	} else {
		fmt.Println(data1["name"])
	}

	data2 := NewMap("Dono")

	if data2 == nil {
		fmt.Println("Data masih kosong")
	} else {
		fmt.Println(data2["name"])
	}
}