Writing a basic HTTP server is easy using the `net/http` package.

A fundamental concept in `net/http` servers is *handlers*.
A handler is an object implementing the `http.Handler` interface.
A common way to write a handler is by using the `http.HandlerFunc` adapter on functions with the appropriate signature.

Functions serving as handlers take a `http.ResponseWriter` and a `http.Request` as arguments.
The response writer is used to fill in the HTTP response.