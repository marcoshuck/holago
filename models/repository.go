package main

import "fmt"

type Repository interface {
	Do(model Model) (Model, error)
}

type userRepository struct {}

func (r userRepository) Do(model Model) (Model, error) {
	var p Person
	err := p.Parse(model)
	if err != nil {
		return nil, err
	}
	fmt.Println("Name:", p.Name, "| Email:", p.Email)
	return &p, nil
}

func NewUserRepository() Repository {
	return &userRepository{}
}
