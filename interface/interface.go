package main

import (
	"fmt"
)

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(totalArea(&c, &r))
}

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
}

type Circle struct {
    x, y, r float64
}

func (c *Circle) area() float64 {
    return 111111111111111
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {

    return 222222222222
}


// Интерфейсы также могут быть использованы в качестве полей:
type MultiShape struct {
    shapes []Shape
}

func (m *MultiShape) area() float64 {
    var area float64
    for _, s := range m.shapes {
        area += s.area()
    }
    return area
}

