package main

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodBye := getGoodBye
	result := goodBye("Nico")
	println(result)
}
