package main

import "fmt"

func deferred() {
	fmt.Println("Yes")
}

func main() {
	defer deferred()
	if true {
		fmt.Println("Do you see me?")
		return
	}
	return
}
