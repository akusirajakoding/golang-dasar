package main

import "fmt"

func main() {
	 names := []string{"Bruce", "Wayne", "Batman", "Alfred", "Robin", "Oracle"}

	 slice1 := names[4:6]
	 slice2 := names[0:3]

	 fmt.Println(names)
	 fmt.Println(slice1)
	 fmt.Println(slice2)
	 fmt.Println(len(slice1))
	 fmt.Println(cap(slice1))

	 slice3 := append(slice1, "Batgirl")
	 slice3 [0] = "Nightwing"
	 slice3 [1] = "Robin"

	 fmt.Println(names)
	 fmt.Println(slice1)
	 fmt.Println(slice2)
	 fmt.Println(slice3)

	 slice4 := make([]string, 7, 7)
	 copy(slice4, names)
	 fmt.Println(slice4)
	 slice4 = append(slice4, "Jason")
	 fmt.Println(slice4)
	 slice4 = append(slice4, "Tim")
	 fmt.Println(slice4)
	 fmt.Println(len(slice4))
	 fmt.Println(cap(slice4))
	}