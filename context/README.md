HTTP servers are useful for demonstrating the usage of context.Context for controlling cancellation.
A Context carries deadlines, cancellation signals, and other request-scoped values across API boundaries and goroutines.

A `context.Context` is created for each request by the `net/http` machinery, and is available with the `Context()` method.
