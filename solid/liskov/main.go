package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Line struct {
	Length float64
}

func (c Line) Area() float64 {
	return 0
}
func PrintArea(s Shape) {
	fmt.Printf("Area : %.2f\n", s.Area())
}
func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}
	line := Line{Length: 7}
	PrintArea(rect)   // Outputs: Area: 50.00
	PrintArea(circle) // Outputs: Area: 153.94
	PrintArea(line)   // Outputs: Area: 0
}
