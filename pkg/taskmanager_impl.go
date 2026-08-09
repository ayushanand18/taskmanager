package main

import (
	"context"
	"sync"

	"github.com/ayushanand18/taskmanager/pkg/types"
)

var _taskManagerInstance TaskManager
var _taskManagerLock sync.Mutex

type taskmanager struct {
	isInitialised bool

	maxNumThreads    uint32
	minNumThreads    uint32
	maxThreadsPerJob uint32
}

func (t *taskmanager) Submit(ctx context.Context, primaryKey string, executeFn types.ExecuteFn) (types.Future, error) {
	return nil, nil
}

func (t *taskmanager) ResolveAll(ctx context.Context, promises ...types.Future) error {
	return nil
}

func (t *taskmanager) Cancel(ctx context.Context, task types.Task) error {
	return nil
}

func (t *taskmanager) Init(ctx context.Context, option types.InitOptions) error {
	return nil
}

func (t *taskmanager) CancelAll(ctx context.Context) error {
	return nil
}

// Close
// CancelAll tasks and free up resources
func (t *taskmanager) Close() {
	t.CancelAll(context.Background())
}

func newTaskManager() TaskManager {
	return &taskmanager{}
}

func GetTaskManager() TaskManager {
	_taskManagerLock.Lock()
	defer _taskManagerLock.Unlock()

	if _taskManagerInstance == nil {
		_taskManagerInstance = newTaskManager()
	}

	return _taskManagerInstance
}
