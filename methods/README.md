Go supports *methods* defined on struct types.

Go automatically handles conversion between values and pointers for method calls.
You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
