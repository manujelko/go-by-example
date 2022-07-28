Go offers built-in support for creating dynamic content or showing customized output to the user with the `text/template` package.
A sibling package named `html/template` provides the same API but has additional security features and should be used for generating HTML.

We can create a new template and parse its body from a string.
Templates are a mix of static text and "actions" enclosed in `{{...}}` that are used to dynamically insert content.

Alternatively, we can use the `template.Must` function to panic in case `Parse` returns an error.
This is especially useful for templates initialized in the global scope.

By "executing" the template we generate its text with specific values for its actions.
The `{{.}}` action is replaced by the value passed as a parameter to `Execute`.

If the data is a struct, we can use the `{{.FieldName}}` action to access its fields.
The fields should be exported to be accessible when a template is executing.

The same applies to maps; with maps there is no restriction on the case of key names.