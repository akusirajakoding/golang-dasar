package main

import "fmt"

func main() {
	batman := map[string]string{
		"name":       "Bruce Wayne",
		"alias":      "Batman",
		"city":	   "Gotham",
		"sidekick":   "Robin",
		"butler":     "Alfred",
		"techExpert": "Oracle",
	}
	fmt.Println(batman)
	fmt.Println(len(batman))
	fmt.Println(batman["alias"])

	fmt.Println(" ")

	robin := make(map[string]string)
	robin["name"] = "Dick Grayson"
	robin["alias"] = "Robin"
	robin["city"] = "Bludheaven"
	robin["sidekick"] = "Batman"
	robin["butler"] = "Alfred"

	delete(robin, "sidekick")
	fmt.Println(robin)
	robin["butler"] = "None"
	fmt.Println(robin)
}