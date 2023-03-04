package main

func main() {
	counter := 1

	for counter <= 10 {
		println(counter)
		counter++
	}

	for counter := 1; counter <= 10; counter++ {
		println(counter)
	}

	names := []string{"Eko", "Kurniawan", "Khannedy"}
	for index, name := range names {
		println(index, name)
	}

	for _, name := range names {
		println(name)
	}

	for index := range names {
		println(index)
	}

	for index := 0; index < len(names); index++ {
		println(names[index])
	}

}
