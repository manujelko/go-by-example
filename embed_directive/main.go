package main

// Import the `embed` package; if you don't use any exported identifiers from
// this package, you can do a blan import with `_ "embed"`.
import "embed"

// `embed` directives accept paths relative to the directory containing the Go source file.
// This directive embed the contents of the file into the `string` variable immediately following it.
//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

// We can also embed multiple files or even folders with wildcards.
// This uses a variable of the `embed.FS` type,
// which implements a simple virtual file system.
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
