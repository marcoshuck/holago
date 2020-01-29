package main

import "fmt"

func createEntity() interface{} {
	return Entity{
		FirstName: "Marcos",
		LastName: "Huck",
	}
}

func preWorker(e interface{}) (interface{}, bool) {
	r, ok := e.(Entity)
	return r, ok
}

func process(e interface{}) interface{} {
	if entity, ok := e.(Entity); ok {
		return fmt.Sprintf("%s %s", entity.FirstName, entity.LastName)
	}
	return nil
}

func postWorker(e interface{}) (interface{}, bool) {
	r, ok := e.(string)
	return r, ok
}

func createStage() *Stage {
	return &Stage{
		PreHook:  preWorker,
		Process:  process,
		PostHook: postWorker,
		Name:     "Example",
	}
}

func main() {
	fmt.Println("Starting application")
	entity := createEntity()

	stage := createStage()

	result, ok := stage.Run(entity)
	if !ok {
		return
	}

	fmt.Println(fmt.Sprintf("%s", result))
}
