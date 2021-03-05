package main

var (
	AWS   string = "aws"
	GCP   string = "gcp"
	Azure string = "azure"
)

func main() {
	inputs := []string{AWS, GCP, GCP, Azure, AWS}
	f := NewMachineFactory()

	managers := make(map[int]MachinesManager)
	for i, in := range inputs {
		managers[i] = f(in)
	}

	for _, m := range managers {
		m.Create()
	}
}
