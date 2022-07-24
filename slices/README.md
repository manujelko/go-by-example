*Slices* are a key data type in Go, giving a more powerful interface to sequences than arrays.

Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
To create an empty slice with non-zero length, use the built-in `make`.

`len` returns the length of the slice as expected.

In addition to these basic operations, slices support several more that make them richer than arrays.
One is the built-in append, which returns a slice containing one or more new values.
Note that we need to accept a return value from append as we might get a new slice value.

Slices can also be `copy`'d.

Slices support a "slice" operator with the sytax `slice[low:high]`.

We can declare an initialize a variable for a slice in a single line.

Slices can be composed into multi-dimensional data structures.
