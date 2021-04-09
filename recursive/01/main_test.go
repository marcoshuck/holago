package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecimalToBinary(t *testing.T) {
	// 0 -> 0
	// 1 -> 1
	// 10 -> 2
	assert.Equal(t, []int{0}, ConvertToBinaryRecursive(0, []int{}))
	assert.Equal(t, []int{1}, ConvertToBinaryRecursive(1, []int{}))
	assert.Equal(t, []int{1, 0}, ConvertToBinaryRecursive(2, []int{}))
	assert.Equal(t, []int{1, 1}, ConvertToBinaryRecursive(3, []int{}))
	assert.Equal(t, []int{1, 0, 0}, ConvertToBinaryRecursive(4, []int{}))
	assert.Equal(t, []int{1, 0, 1}, ConvertToBinaryRecursive(5, []int{}))
	assert.Equal(t, []int{1, 1, 0}, ConvertToBinaryRecursive(6, []int{}))
	assert.Equal(t, []int{1, 1, 1}, ConvertToBinaryRecursive(7, []int{}))

	assert.Equal(t, []int{0}, ConvertToBinaryLoop(0))
	assert.Equal(t, []int{1}, ConvertToBinaryLoop(1))
	assert.Equal(t, []int{1, 0}, ConvertToBinaryLoop(2))
	assert.Equal(t, []int{1, 1}, ConvertToBinaryLoop(3))
	assert.Equal(t, []int{1, 0, 0}, ConvertToBinaryLoop(4))
	assert.Equal(t, []int{1, 0, 1}, ConvertToBinaryLoop(5))
	assert.Equal(t, []int{1, 1, 0}, ConvertToBinaryLoop(6))
	assert.Equal(t, []int{1, 1, 1}, ConvertToBinaryLoop(7))
}

func ConvertToBinaryRecursive(decimal int, binary []int) []int {
	// Si el número es 0 o 1, ya se que es el ultimo numero, entonces lo agrego al slice y retorno.
	if decimal < 2 {
		return append(binary, decimal)
	}
	// Calculo el siguiente valor llamando recursivamente a la misma funcion, pero ojo, el valor de entrada esta divido por 2.
	binary = ConvertToBinaryRecursive(decimal/2, binary)

	// Al terminar de hacer todas las llamadas, agrego el resultado. Esto se puso después que la llamada a la funcion recursiva porque sino
	// el numero salia dado vuelta.
	return append(binary, decimal%2)
}

func ConvertToBinaryLoop(decimal int) []int {
	var binary []int
	for decimal >= 2 {
		binary = append([]int{decimal % 2}, binary...)
		decimal = decimal / 2
	}
	binary = append([]int{decimal % 2}, binary...)
	return binary
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertToBinaryLoop(999)
	}
}

func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertToBinaryRecursive(999, []int{})
	}
}
