Unit testing is an important part of writing principled Go programs.
The `testing` package provides the tools we need to write unit tests and the `go test` command runs tests.

For the sake of demonstration, this code is in package main, but it could be any package.
Testing code typically fives in the same package as the code it tests.

Typically, the code we're testing would be in a source file name something like `intutils.go`, and the test file for it would then be named `intutils_test.go`.

A test is created by writing a function with a name beginning with `Test`.

`t.Error*` will report test failures but continue executing the test.
`t.Fatal*` will report test failures and stop the test immediately.
