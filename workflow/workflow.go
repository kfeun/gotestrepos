package main

import (
	"fmt"
)

func main() {

	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i += 2
	// }

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	i := 3

	switch i {
	case 0:
		fmt.Println("ZERO")
	case 1:
		fmt.Println("ONE")
	case 2:
		fmt.Println("TWO")
	case 3:
		fmt.Println("THREE")
	default:
		fmt.Println("Unknown Number")
	}
}
