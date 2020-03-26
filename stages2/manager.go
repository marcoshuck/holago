package main

import (
	"context"
	"fmt"
)

type Manager func(ctx context.Context, entity interface{}, stages []IStage) (interface{}, error)



func ApplicationManager(ctx context.Context, entity interface{}, stages []IStage) (interface{}, error) {
	var result interface{}
	var err error
	for i, stage := range stages {
		fmt.Println(fmt.Sprintf("Running the [%d] stage -- Name: %s", i, stage.GetName()))
		result, err = stage.Run(entity)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

