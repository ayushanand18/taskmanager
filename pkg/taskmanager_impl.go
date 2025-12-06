package main

import (
	"context"

	"github.com/ayushanand18/taskmanager/pkg/types"
)

type taskManager struct {
	isInitialised bool

	maxNumThreads uint32
	minNumThreads uint32
	maxThreadsPerJob uint32
}

func (t *taskManager) Submit (ctx context.Context, executeFn types.ExecuteFn) (types.Promise, error) {
	return nil, nil
}

func (t *taskManager) ResolveAll (ctx context.Context, promises ...types.Promise) error {
	return nil
}
