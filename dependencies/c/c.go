package c

import (
	"fmt"
	"github.com/marcoshuck/holago/dependencies/a"
)

type C interface {
	DoC()
	SetDep(dep a.A)
}

type c struct {
	dep a.A
}

func (c *c) SetDep(dep a.A) {
	c.dep = dep
}

func (c c) DoC() {
	fmt.Println("Hello from C")
	c.dep.Finish()
}

func NewC() C {
	var value C
	value = &c{}
	return value
}