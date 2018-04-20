package main

import "fmt"

func main() {

	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	// Аргументы (входы) определяются так: имя тип, имя тип, ….
	x, y := f()
	fmt.Println(x, y)
	fmt.Println(add(1, 2, 3))

	newAdd := func(x, y int) int {
		return x + y
	}
	fmt.Println(newAdd(1, 1))

	// замыкание
	z := 0
	increment := func() int {
		z++
		return z
	}
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println(factorial(15))
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func f() (int, int) {
	return 5, 6
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// рекурсия
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}
