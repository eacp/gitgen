# Gitgen
Generate license and gitignore files from Go without an internet connection. It also has a convenience CLI, 
but can be used as a library as well.

This is a library to generate .gitignore and LICENSE files directly from Go, locally and without the need for an internet connection 
or external APIs.  It takes advantage of the `embed` package aded in Go 1.16 to embed multiple .gitignore and license templates 
in the library, and provides functions to interact with them.

With this package you can directly get the text of the files as a string, use the license templates with your parameters to directly 
generate `LICENSE` files with the required year and author, and even write to any `io.Writer`, such as files or `http.ResponseWriter`

## Examples

### Get the text of a `.gitignore` template

```go

nodeGitIgnore := gitgen.GetIgnoreText("Node")
println(nodeGitIgnore) // prints the entire .gitignore file for Node

```


### Write a `.gitignore` to a file

```go
f, err := os.Create(".gitignore")
defer f.Close()

// Do something with the error

_, err = gitgen.WriteIgnore("Java", f)

// Do something with the error

```


### Get the text of a `LICENSE` template

```go

mitLic := gitgen.GetLicenseText("mit")
println(mitLic) // prints the entire LICENSE

```


### Write a `LICENSE` to a file

```go
f, err := os.Create("LICENSE")
defer f.Close()

// Do something with the error

_, err = gitgen.WriteLicense("mit", f)

// Do something with the error

```

### Get the text of a `LICENSE` template, but with YEAR and NAME parameters

```go

mitLic := gitgen.GetLicWithParams("mit", "eacp", "2021")
println(mitLic) // prints the entire LICENSE with the name and the year

```


### Write a `LICENSE` to a file with name and year

```go
f, err := os.Create("LICENSE")
defer f.Close()

// Do something with the error

_, err = gitgen.WriteLicWithParams("mit", "eacp", "2021", f)

// Do something with the error

```

## Testing

You can test the package by using the `go test` command. It has been tested on Windows and Linux. Mac testing is pending. 

# CLI

I also made a CLI around the library for conveniance purposes. Go to the cli folder, execute it and use `cli help` for more information. 
The cli can also be tested by using the `go test`command
