package main

import "fmt"

// Component is the component we want to use inside a job
type Component interface {
	Do(input interface{})
}

// component is a Component implementation.
type component struct{}

// Do does something with the given input. In this case it appends the input to the Hello World string and
// shows it in the std output.
func (component) Do(input interface{}) {
	fmt.Println("Hello world!", input)
}

// NewComponent initializes a new Component.
func NewComponent() Component {
	return &component{}
}

// main represents an action creation.
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
