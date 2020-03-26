package main

import (
	"context"
	"errors"
	"fmt"
)

type IApplication interface {
	Prepare() context.Context
	Start(ctx context.Context)
	Finish(ctx context.Context)
	HasFinished() bool
	Panic()
}

type Application struct {
	state int
	finished bool
	ApplicationManager Manager
}

func (app *Application) Prepare() context.Context {
	fmt.Println("[APP] Preparing application.")
	ctx := context.Background()
	app.state = 0
	app.finished = false
	app.ApplicationManager = ApplicationManager
	fmt.Println("[APP] Application has been successfully prepared.")
	return ctx
}

func (app *Application) Start(ctx context.Context) {
	fmt.Println("[APP] Starting application.")
	stages := app.setup()
	value, err := app.ApplicationManager(ctx, app.state, stages)
	if err != nil {
		fmt.Println("[APP] There's been an error during application runtime.")
		app.Finish(ctx)
		panic(err)
	}
	result, ok := value.(int)
	if !ok {
		fmt.Println("[APP] The result is invalid.")
		app.Panic()
		app.Finish(ctx)
	}
	fmt.Println("[APP] Result:", result)
	app.Finish(ctx)

}

func (app *Application) setup() ([]IStage) {
	stages := []Stage{}
}

func (app *Application) Finish(ctx context.Context) {
	fmt.Println("[APP] The application has finished")
	app.finished = true
}

func (app Application) Panic() {
	panic(errors.New("fatal error"))
}

func (app Application) HasFinished() bool {
	return app.finished
}