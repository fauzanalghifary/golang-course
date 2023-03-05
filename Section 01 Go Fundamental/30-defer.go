package main

func logging() {
	println("finish invoke function")
}

func runApplication(value int) {
	defer logging()
	println("application running")
	result := 10 / value
	println("result", result)
}

func main() {
	runApplication(0)
}
