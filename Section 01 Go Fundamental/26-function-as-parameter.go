package main

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) string {
	nameFiltered := filter(name)
	return "Hello " + nameFiltered
}

func spamFilter(name string) string {
	if name == "Sh*t" {
		return "..."
	} else {
		return name
	}
}

func main() {
	result := sayHelloWithFilter("Sh*t", spamFilter)
	println(result)
}
