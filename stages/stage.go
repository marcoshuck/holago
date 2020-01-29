package main

import "fmt"

type Stage struct {
	PreHook func(interface{}) (interface{}, bool)
	Process func(interface{}) interface{}
	PostHook func(interface{}) (interface{}, bool)
	Name string
}

func (s *Stage) Run(e interface{})  (interface{}, bool) {
	fmt.Println(fmt.Sprintf("Starting to run stage: %s", s.Name))

	fmt.Println(fmt.Sprintf("Running prehook in stage: %s", s.Name))
	if r, ok := s.PreHook(e); !ok {
		return r, ok
	}

	fmt.Println(fmt.Sprintf("Running process in stage: %s", s.Name))
	result := s.Process(e)


	fmt.Println(fmt.Sprintf("Running posthook in stage: %s", s.Name))
	if r, ok := s.PostHook(result); !ok {
		return r, ok
	}

	return result, true
}