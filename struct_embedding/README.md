Go supports *embedding* of structs and interfaces to express more seamless composition of types.
This is not to be confused with //go:embed which is a go directive introduced in Go version 1.16+ to embed files and folders into the application binary.

Embedding structs with methods may be used to bestow interface implementations onto other structs.