Maps are Go's built-in associative data type.

To create an empty map use the built-in make.

Set key/value pairs using typical `name[key] = value` syntax.

Get a value for the key with `name[key]`.

The builtin `len` returns the number of key/value pairs in the map.

The builtin `delete` removes key/value pairs from a map.

The optional second return value when getting a value from a map indicates if a key was present in the map.
This can be used to disambiguate between missing keys and keys with zero value like 0.
