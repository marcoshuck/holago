package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInvalidStruct(t *testing.T) {
	in := Cat{
		Name:  "Ginny",
		Sound: "Meow!",
	}
	var p Person
	err := p.Parse(&in)
	assert.Error(t, err)
}

func TestParseValidStruct(t *testing.T) {
	in := Person{
		Name: "Marcos",
		Email: "marcos@huck.com.ar",
	}

	var p Person
	err := p.Parse(&in)
	assert.NoError(t, err)
}

func TestUserRepository(t *testing.T) {
	var repo Repository
	repo = NewUserRepository()
	
	in := Person{
		Name: "Marcos",
		Email: "marcos@huck.com.ar",
	}
	
	p, err := repo.Do(&in)
	assert.NoError(t, err)
	assert.Equal(t, &in, p)
}