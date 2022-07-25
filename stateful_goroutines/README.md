In this example our state will be owned by a single goroutine.
This will guarantee that the data is never corrupted with concurrent access.
In order to read or write that state, other goroutines will send messages to the owning goroutine and receive corresponding replies.
These readOp and writeOp structs encapsulates those requests and a way for the owning goroutine to respond.
