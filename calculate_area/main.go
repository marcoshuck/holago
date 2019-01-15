package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radious float64
}

func (sq square) area() float64 {
	return sq.side * sq.side
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radious, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{5}
	c := circle{2}
	info(sq)
	info(c)
}
