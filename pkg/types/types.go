package types

import "context"

type ExecuteFn func(context.Context) error 

type Promise interface {
	Cancel(context.Context) error
	Resolve(context.Context) error
}

type InitOptions struct {
	MaxNumThreads uint32
	MinNumThreads uint32
	MaxThreadsPerJob uint32
}
