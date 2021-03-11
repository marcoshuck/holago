package main

import "fmt"

func main() {
	fmt.Println(SquaresOf(Integers(25)))
}

func SquaresOf(list []int) []int {
	for i, value := range list {
		list[i] = value * value
	}
	return list
}

func Integers(n int) []int {
	list := make([]int, n)
	for i := 1; i <= n; i++ {
		list[i-1] = i
	}
	return list
}
