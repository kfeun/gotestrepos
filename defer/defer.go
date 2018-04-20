package main

import (
	"fmt"
	// "os"
)

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func main() {
	// позволяет отложить вызов указанной функции до тех пор, пока не завершится текущая.
	defer second()
	first()

	// f, _ := os.Open(filename)
	// defer f.Close()

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
