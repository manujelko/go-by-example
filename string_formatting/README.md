Go offers excellent support for string formatting int the printf tradition.

If the value is a struct, the `%+v` will include the structs field's names.

The `%#v` variants prints a Go syntax representation of the value, i.e. the source code snippet that would produce that value.

To print the type of a value, use `%T`.
