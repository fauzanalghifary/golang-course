package main

import (
	"errors"
	"fmt"
)

func Divider(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("divider can't be 0")
	} else {
		result := value / divider
		return result, nil
	}
}

func main() {
	var errorExample error = errors.New("Ups something wrong")
	fmt.Println(errorExample.Error())

	result, err := Divider(10, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
