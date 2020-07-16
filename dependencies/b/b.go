package b

import (
	"fmt"
	"github.com/marcoshuck/holago/dependencies/c"
)

type B interface {
	DoB()
	SetDep(dep c.C)
}

type b struct {
	dep c.C
}

func (b *b) SetDep(dep c.C) {
	b.dep = dep
}

func (b *b) DoB() {
	fmt.Println("Hello from B")
	b.dep.DoC()
}

func NewB() B {
	var value B
	value = &b{}
	return value
}