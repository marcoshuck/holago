package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buff bytes.Buffer

	input := []byte("test")

	buff.Read(input)

	var test []byte
	io.ReadFull(&buff, test)

	fmt.Println(string(test))
}
