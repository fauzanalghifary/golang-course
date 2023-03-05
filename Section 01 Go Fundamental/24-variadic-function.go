package main

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 20, 30)
	println(total)

	//	Slice parameter
	slice := []int{10, 20, 30}
	total = sumAll(slice...)
	println(total)
}
