package main 

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("berhenti di perulangan ke-", i)
			break
	}
		fmt.Println("perulangan ke-", i)
	}
}