# taskmanager
A simple taskmanager implementation

## design
Runtime
```go
Submit(primary_key) -
  submits a task to the shared threadpool. 
  primary_key -> rate limit + dedicated pools

ResolveAll([...promises]) -> are all promises resolved

Resolve(promise) -> blocking wait for the promise to resolve
```

Initialisation
```go
Options{
	MaxNUmberOfThreads: uint32,
	MinNumberOfThreads: uint32,
	PerJobMaxNumThreads: uint32
}
Init(context.Context, Options)

```
