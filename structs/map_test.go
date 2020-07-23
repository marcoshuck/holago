package structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type user struct {
	Name *string `json:"name,omitempty"`
	Password *int `json:"password,omitempty"`
}

func TestToMap_OK(t *testing.T) {
	name := "marcos"
	password := 1234
	u := user{
		Name:     &name,
		Password: &password,
	}

	m, err := toMap(u)
	assert.NoError(t, err)
	assert.Equal(t, name, m["name"])
	assert.Equal(t, password, m["password"])
}

func TestToMap_Invalid(t *testing.T) {
	_, err := toMap(1234)
	assert.Error(t, err)
}

func TestToMap_MissingFields(t *testing.T) {
	name := "marcos"
	u := user{
		Name:     &name,
		Password: nil,
	}
	m, err := toMap(u)
	assert.NoError(t, err)
	assert.Equal(t, name, m["name"])
	assert.Equal(t, nil, m["password"])
}

func TestToMap_ZeroValues(t *testing.T) {
	name := ""
	password := 0
	u := user{
		Name:     &name,
		Password: &password,
	}
	m, err := toMap(u)
	assert.NoError(t, err)
	assert.Equal(t, name, m["name"])
	assert.Equal(t, password, m["password"])
}

func BenchmarkToMap(b *testing.B) {
	name := "marcos"
	password := 1234
	u := user{
		Name:     &name,
		Password: &password,
	}
	for i := 0; i < b.N; i++ {
		_, err := toMap(u)
		if err != nil {
			b.Error(err)
		}
	}
}