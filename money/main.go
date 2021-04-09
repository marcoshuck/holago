package main

import "fmt"

func main() {
	v := 0.2
	for i := 0; i < 100; i++ {
		v += 0.2
	}
	fmt.Printf("%.12f", v)
}
