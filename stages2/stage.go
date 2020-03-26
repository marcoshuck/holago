package main

import (
	"context"
	"fmt"
)

type IStage interface {
	GetTag() string
	GetName() string
	Run(entity interface{}) (interface{}, error)
}

type Stage struct {
	observers []IObserver
	tag string
	name string
	/////////////////////////////////////////////////////////////////////////////
	// Init defines an event that will be triggered at the beginning of the stage
	Init IEvent
	AfterInit IHook
	BeforeProcess IHook
	Process IEvent
	AfterProcess IHook
	BeforeFinish IHook
	// Finish defines an event that will be triggered at the end of the stage
	Finish IEvent
}

func (st Stage) GetTag() string {
	return st.tag
}

func (st Stage) GetName() string {
	return st.name
}

func (st *Stage) Run(entity interface{}) (interface{}, error) {

}

func (st *Stage) Add(ctx context.Context, observer IObserver) {
	st.observers = append(st.observers, observer)
}

func (st *Stage) Remove(ctx context.Context, observer IObserver) {
	for i, ob := range st.observers {
		if ob == observer {
			st.observers = append(st.observers[:i], st.observers[i+1:]...)
		}
	}
}

func (st *Stage) Notify(ctx context.Context, status string) {
	for i, observer := range st.observers {
		fmt.Println(fmt.Sprintf("[%s] [%d] %s: %s", st.GetTag(), i,  st.GetName(), status))
		observer.Update(ctx)
	}
}