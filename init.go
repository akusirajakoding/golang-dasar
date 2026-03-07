package main

import (
	"fmt"
	"golang-dasar/database"
	_ "golang-dasar/internal" // gunakan (_) untuk eksekusi initnya doang
)

func main() {
	fmt.Println(database.GetDatabase())
}
