package main

import (
	"context"

	"github.com/ayushanand18/taskmanager/pkg/types"
)

// TaskManager - handle multitudes of task in a CPU constrained environment
// scales from a MinThreadCount to MaxThreadCount based on load
// when more tasks add up, it does not spawn more threads than MaxThreadCount,
// instead queues up tasks until threads are free
type TaskManager interface {
	// Submit a task to the task manager, it will write to a response channel once done
	// returns taskContext, which can be cancelled to stop the task
	// 		   error -> if task could not be submitted
	Submit(ctx context.Context, primaryKey string, executeFn types.ExecuteFn) (types.Future, error)
	// Cancel a task previously submitted to the task manager
	Cancel(ctx context.Context, task types.Task) error
	// CancelAll cancels all tasks in the task manager
	CancelAll(ctx context.Context) error
	// ResolveAll - makes a blocking call until all promises are resolved
	ResolveAll(ctx context.Context, promises ...types.Future) error
	// Init - initialization of pool
	Init(ctx context.Context, option types.InitOptions) error
	// Close the task manager and free up resources
	Close()
}
