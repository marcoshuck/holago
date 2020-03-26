package main

import "context"

type IEvent func(context.Context, interface{}) (interface{}, error)