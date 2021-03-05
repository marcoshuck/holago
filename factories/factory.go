package main

func NewMachineFactory() func(config string) MachinesManager {
	return func(config string) MachinesManager {
		switch config {
		case AWS:
			return NewAWSMachineManager(config)
		case GCP:
			return NewGCPMachineManager(config)
		case Azure:
			return NewAzureMachineManager(config)
		default:
			panic("machine factory: invalid config provided")
		}
	}
}
