# Go Workshop

This repository contains the Go code workshop to help colleagues
get started with writing Go code.
In the workshop we touch upon a few basic Go constructs and the Gin framework
that allows us to easily build an API with only a few lines of code.

- [Go Workshop](#go-workshop)
  - [Prerequisites](#prerequisites)
  - [Getting started](#getting-started)
    - [Gin can not be found](#gin-can-not-be-found)
  - [Material](#material)
  - [Useful commands](#useful-commands)
  - [Disclaimer](#disclaimer)

## Prerequisites

For this workshop you're required to have [Go installed](https://golang.org/doc/install).
You also might want to have Make installed to utilize the makefile commands.

During the workshop the [GoLand IDE](https://www.jetbrains.com/go/) will be used by the host,
but it's perfectly OK to use Visual Studio Code or any other IDE that can read Go code.

## Getting started

1. Clone this repository
1. Make sure gin is installed: `go get github.com/codegangsta/gin`
1. Run `make t` to run the tests
1. Run `make r` to start the application or use `gin --appPort 8080 -i`
1. Visit the following pages: 
   - Basics: [localhost:8080/basics/hello-world](http://localhost:8080/basics/hello-world)
   - API: [localhost:8080/api/people](http://localhost:8080/api/people)

### Gin can not be found

If your shell is unable to find `gin`, it means your `PATH` environment variable
does not include the directory where go installs binaries.
Go to your home directory (`cd ~`) and look for the `go` directory, inside it there
should be a `bin` directory where the binaries are installed.

Learn how to add this directory to your `PATH` here:

- [Windows](https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/)
- [MacOS](https://www.architectryan.com/2012/10/02/add-to-the-path-on-mac-os-x-mountain-lion/)
- [Linux (bash)](https://docs.oracle.com/cd/E19062-01/sun.mgmt.ctr36/819-5418/gaznb/index.html)

## Material

The workshop is divided up in these 3 sections.

| Section | Reference material | Exercises | Expected time | Description
| ------- | ---- | ---- | ------------- | ------------
| Basics  | [Link](api/basics/reference-material.md) | [Link](api/basics/exercises.md) | 2 hours | Going over the basics of the Go language such as syntax, differences with OOP languages and exception handling
| API     | [Link](api/workshop-api/reference-material.md) | [Link](api/workshop-api/exercises.md) | 2 hours | Finishing the API with a few more routes like a POST route
| Useful patterns | [Link](api/patterns/reference-material.md)| [Link](api/patterns/exercises.md) | 1 hour, 30 minutes | Providing a list of useful patterns to improve code testing and quality

## Useful commands

Some useful commands include:

- `go test`: Runs your tests in the current package (We use `make t` for this)
- `go run`: Runs the application locally (but we're using [gin](https://github.com/codegangsta/gin) for this workshop)
- `go build`: Builds the application to a single binary called `P02012-Go-Workshop`
  
---

- `go fmt`: Formats your Go code
- `go get`: Downloads and installs go packages (dependencies)

For more useful commands [see this article](https://www.ubuntupit.com/go-command-examples-for-aspiring-golang-developers/)

## Disclaimer

Although we're experienced with Go, we're not experts.
If you discover new/better ways to go about certain topics please
let us know!
