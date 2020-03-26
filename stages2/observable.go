package main

import "context"

type IObservable interface {
	Add(ctx context.Context, observer IObserver)
	Remove(ctx context.Context, observer IObserver)
	Notify(ctx context.Context)
}