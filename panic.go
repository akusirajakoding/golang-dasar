package main

import "fmt"

func endApp() {
	fmt.Println("endApp")
	message := recover()
	fmt.Println("ups terjadi panic,", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("panic error")
	}
	// jika meletakkan recover di bawah panic yg error maka tidak akan terjadi recover
	// karena jika terjadi panic kode di bawahnya tidak akan dilanjutkan kecuali defer
	// oleh karena itu letakkan recover di dalam defer
	//message := recover()
	//fmt.Println("ups terjadi panic,", message)
}

func main() {
	runApp(true)
}
