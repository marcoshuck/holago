package main

import "fmt"

type MachinesManager interface {
	Create()
	Destroy()
}

type fakeMachinesManager struct {
	provider string
}

func (f *fakeMachinesManager) Create() {
	fmt.Println("Creating machines with", f.provider)
}

func (f *fakeMachinesManager) Destroy() {
	fmt.Println("Destroying machines with", f.provider)
}

func NewAzureMachineManager(config string) MachinesManager {
	return &fakeMachinesManager{
		provider: config,
	}
}

func NewGCPMachineManager(config string) MachinesManager {
	return &fakeMachinesManager{
		provider: config,
	}
}

func NewAWSMachineManager(config string) MachinesManager {
	return &fakeMachinesManager{
		provider: config,
	}
}

type StorageManager interface {
	Read()
	Save()
}
