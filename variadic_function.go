package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sumAll(1, 2, 3)
	fmt.Println(total)

	numbers := []int{1, 2, 3, 4}
	fmt.Println(sumAll(numbers...))

	jumlahSemuaAngka := sumAll(23, 83, 90, 42)
	fmt.Println(jumlahSemuaAngka)

	sumSlice := []int{32, 42, 80, 54}
	fmt.Println(sumAll(sumSlice...))
}
