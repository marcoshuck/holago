package main

import (
	"github.com/marcoshuck/holago/dependencies/a"
	"github.com/marcoshuck/holago/dependencies/b"
	"github.com/marcoshuck/holago/dependencies/c"
)

func main() {
	app := a.NewA()
	depB := b.NewB()
	depC := c.NewC()

	app.SetDep(depB)
	depB.SetDep(depC)
	depC.SetDep(app)

	app.Start()
}