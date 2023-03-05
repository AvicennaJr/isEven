package main

import (
	"fmt"

	iseven "github.com/AvicennaJr/isEven"
)

func main() {
	n := 4
	if iseven.IsEven(n) {
		fmt.Println("It is even")
	} else {
		fmt.Println("It is not even")
	}
}
