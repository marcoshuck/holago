package main

import "errors"

type Model interface {
	Parse(model Model) error

}

type Person struct {
	Name string
	Email string
}

func (p *Person) Parse(model Model) error {
	result, ok := model.(*Person)
	if !ok {
		return errors.New("invalid Person model")
	}
	*p = *result
	return nil
}

type Cat struct {
	Name string
	Sound string
}

func (c *Cat) Parse(model Model) error {
	result, ok := model.(*Cat)
	if !ok {
		return errors.New("invalid Cat model")
	}
	*c = *result
	return nil
}