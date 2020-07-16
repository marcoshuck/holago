package a

import (
	"fmt"
	"github.com/marcoshuck/holago/dependencies/b"
)

type A interface {
	Start()
	Finish()
	SetDep(dep b.B)
}

type a struct {
	dep b.B
}

func (a *a) SetDep(dep b.B) {
	a.dep = dep
}

func (a *a) Start() {
	fmt.Println("Start from A")
	a.dep.DoB()
}

func (a *a) Finish() {
	fmt.Println("Finish from A")
}

func NewA() A {
	var value A
	value = &a{}
	return value
}