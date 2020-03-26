package main

import "context"

type IHook func(context.Context, interface{}) (interface{}, error)

