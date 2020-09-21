package main

import "fmt"

type Component interface {
	Do(input interface{})
}

type component struct{}

func (component) Do(input interface{}) {
	fmt.Println("Hello world!", input)
}

func NewComponent() Component {
	return &component{}
}

func main() {
	c := NewComponent()
	fn := NewFunc(c)
	j := NewJob(fn)
	j.Do("Marcos!")
}

type Job struct {
	fn func(input interface{}) interface{}
}

func (j *Job) Do(input interface{}) {
	j.fn(input)
}

func NewJob(fn func(input interface{}) interface{}) *Job {
	return &Job{
		fn: fn,
	}
}

func NewFunc(c Component) func(input interface{}) interface{} {
	return func(input interface{}) interface{} {
		c.Do(input)
		return input
	}
}
