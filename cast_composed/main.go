package main

import "fmt"

type A struct {
	ValueA int
}

type B struct {
	ValueB string
}

type C struct {
	A
	B
}

func main() {
	var a interface{}
	a = A{ValueA: 1}
	first, ok := a.(C)
	fmt.Println("First:", first, "|", "OK:", ok)

	var c interface{}
	c = C{
		A: A{
			ValueA: 2,
		},
		B: B{
			ValueB: "test",
		},
	}

	second, ok := c.(B)
	fmt.Println("Second:", second, "|", "OK:", ok)
}