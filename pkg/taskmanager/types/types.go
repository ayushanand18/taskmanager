package types

import "context"

type ExecuteFn func(context.Context) error

type Future interface {
	Cancel(context.Context) error
	Resolve(context.Context) error
}

type InitOptions struct {
	MaxNumThreads    uint32
	MinNumThreads    uint32
	MaxThreadsPerJob uint32
}

type Task struct {
	Func         func(context.Context) error
	ResponseChan chan any
	Request      any
}
