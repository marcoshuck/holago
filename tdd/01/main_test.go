package main

import "testing"

func TestDictionary(t *testing.T) {
	// Input
	const (
		a = "hello"
	)

	// Output
	const (
		b = "a way of greeting people"
	)

	if GetDefinition(map[string]string{
		a: b,
	}, a) != b {
		t.FailNow()
	}
}

func TestDictionaryWhenPassingADifferentWord(t *testing.T) {
	// Input
	const (
		a = "hola"
	)

	// Output
	const (
		b = "una manera de saludar a alguien"
	)

	if GetDefinition(map[string]string{
		a: b,
	}, a) != b {
		t.FailNow()
	}
}

func GetDefinition(data map[string]string, key string) string {
	return data[key]
}
