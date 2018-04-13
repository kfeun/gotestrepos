package main

import "fmt"

var y = "HelloOutside"

const newConst string = "constant"

func main() {
	fmt.Println("Wdwdqw90-9dqd"[0])
	fmt.Println("Hello " + "world")
	fmt.Println(true || false)
	fmt.Println(!false)

	var x string = "Hello\n"
	//string type not needed
	x += "world!"
	fmt.Println(x)

	var y = "HelloInside"
	fmt.Println(y == "Hello")

	z := "Присвоить значение при создании"
	fmt.Println(z)
	fmt.Println(y)
	testFunc()

}

func testFunc() {
	fmt.Println(newConst)

	var (
		a = 1
		b = 2
		c = 3
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
