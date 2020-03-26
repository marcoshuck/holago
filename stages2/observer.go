package main

import "context"

type IObserver interface {
	GetTag() string
	GetName() string
	Update(ctx context.Context)
}