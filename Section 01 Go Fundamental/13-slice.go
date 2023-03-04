package main

import "fmt"

func main() {
	var months = [...]string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}

	var slice1 = months[4:7]
	var slice2 = months[10:]
	var slice3 = months[:3]

	fmt.Println(slice1)
	fmt.Println(cap(slice2))
	fmt.Println(cap(slice3))

	var slice4 = append(slice1, "New")
	fmt.Println(slice4)
	fmt.Println(slice1)

	// make slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "A"
	newSlice[1] = "B"

	fmt.Println(newSlice)

	// copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// delete slice
	// delete slice index 1
	copySlice = append(copySlice[:1], copySlice[2:]...)
	fmt.Println(copySlice)

	//	Slice vs Array
	thisIsArray := [3]int{1, 2, 3}
	thisIsSlice := []int{1, 2, 3}
	fmt.Println(thisIsArray)
	fmt.Println(thisIsSlice)

}
