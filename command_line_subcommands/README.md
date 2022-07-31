Some command-line tools, like the `go` tool or `git` have many *subcommands*, each with its own set of flags.
For example, `go build` and `go get` are two different subcommands of the `go` tool.
The `flag` package lets us easily define simple subcommands that have their own flags.
