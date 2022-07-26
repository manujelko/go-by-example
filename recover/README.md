Go makes it possible to *recover* from a panic by using the `recover` built-in function.
A `recover` can stop a panic from aborting the program and let it continue with execution instead.

Recover must be called within a defered function.
When the enclosing function panics, the defer will activate and a recover call within it will catch the panic.