package main

import "fmt"

func endApp() {
	message := recover()
	fmt.Println(message)
	fmt.Println("end App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
	fmt.Println("App is running")
}

func main() {
	runApp(true)
	fmt.Println("App is still running")
}
