package main 

import "fmt"

func main() {
	nilai := []int{90, 70, 0, 85, 60, 0, 100}
	fmt.Println(nilai)

	for _, value := range nilai {
		if value == 0 {
			continue
		} else if value == 100 {
			fmt.Println("Nilai sempurna ditemukan")
			break
		}
		fmt.Println("Nilai : ", value)
	}
}