package main

type TaskManager interface {
	// Submit - submit a task to shared threadpool
	Submit(primaryKey string) Future
	// ResolveAll - makes a blocking call until all promises are resolved
	ResolveAll(promises Promise...)
	// Init - initialization of pool
	Init(context.Context, option Option)
}


