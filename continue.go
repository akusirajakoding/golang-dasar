package main 

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			// fmt.Println("melewati perulangan ke-", i)
			continue
		}
		fmt.Println("perulangan ke-", i)
	}
}