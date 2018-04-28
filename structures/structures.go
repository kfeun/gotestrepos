package main

import (
	"math"
	"fmt"
)

type Circle struct {
	x float64
	y float64
	r float64
	// x, y, r float64
}

func main() {
	// var c Circle
	// c := new(Circle)
	// c := Circle{x: 0, y: 0, r: 5}
	c := Circle{0, 0, 5}
	c.y = 100
	fmt.Println(c.area())
	p := new(Person)
	p.Talk()
	a := new(Android)
	a.Person.Talk()
	a.Talk()
}

func (c *Circle) area () float64 {
	return math.Pi * c.r*c.r
}

type Person struct {
    Name string
}

type Android struct {
    Person
    Model string
}

func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}
