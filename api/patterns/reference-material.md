# Patterns

Hey welcome to the last section!
By now you should have a working API that allows you to retrieve, filter
and add people to an in-memory dataset.
There are however, a few issues with this code, mainly
its testability.

- [Patterns](#patterns)
  - [Dependency Inversion](#dependency-inversion)
  - [Logging](#logging)
  - [Higher order functions](#higher-order-functions)
  - ['Inheritance' -> Embedding](#inheritance---embedding)

---

## Dependency Inversion

Function `GetPeopleRoute` makes a call to `getPeople`, that means
that during a test, I'm also implicitly testing getPeople right?

What happens if I make calls to external libraries? Or 
if I make use of a database, what happens if I have
a route connect to a database, does that mean I must always
be connected to a database?

In other languages, you'd probably be using a Mocking framework
to prevent making any real calls to a database.
In Go this is not really an option, so we have to
be smart instead.
How do we mock this function? Interfaces.

The `exercises.md` will take you through this journey.

---

## Logging

In the Basics part we used `fmt.Println()` to print messages to the
console, this is useful for make-shift debugging but it is not
actual logging.

To make logging possible, it's useful to use a logger.
Luckily Go has a built-in logger for this.
[Click here to learn more about this package](https://golang.org/pkg/log/).

There are also more packages available, [click here for examples](https://www.datadoghq.com/blog/go-logging/).

---

## Higher order functions

It's also possible to pass functions on to other functions,
allowing you to decouple individual functions and create small
re-usable blocks of code.
These functions are called higher order functions.
Take the following example.

```go
package main

import "strings"

/*
Validation options
*/

// We call it a 'ValidatorFunction' if the input is 'string' and output is 'bool'
type ValidatorFunction func(string) bool

// Signature is (string) => bool
func RequiredVal(input string) bool {
  // Check whether the string is empty
  return strings.Trim(input, " ") == ""
}

// Signature is (string) => bool
func LowercaseVal(input string) bool {
  // Check whether the string is lowercase
  return strings.ToLower(input) == input
}

/**
The function that uses these validation options
*/

// We have the input string and a list of validators we wish to apply
// Remember that these validators are functions that accept 1 input string
// and return a boolean
func ValidateString(input string, validators []ValidatorFunction) bool {
  // Loop through all validators
  for _, validator := range validators {

    // Call the validator function and capture whether the
    // string is valid
    isValid := validator(input)

    // If it's not valid, return false.
    // Otherwise continue executing.
    if !isValid {
      return false
    }
  }

  // Yay! Our string is valid according to the given validators
  return true
}

func main() {
  string1 := "Hello World"
  
  // Check whether the string is not empty,
  // Returns: true
  ValidateString(string1, []ValidatorFunction{RequiredVal})

  // Check whether the string is not empty AND only lowercase,
  // Returns: false
  ValidateString(string1, []ValidatorFunction{RequiredVal, LowercaseVal})
}
```

You can also return functions from functions which in turn can also
return functions ;-)

```go
package main

// Obviously this is a comedic example to illustrate
// how far you can go with these functions.
// Scroll a bit down to get a simpler example
func sum3Numbers(number1 int) func (int) func (int) int {
	return func(number2 int) func (number3 int) int {
		return func (number3 int) int {
			return number1 + number2 + number3
        }
    }
}

// sum2Numbers accepts a number (number1) and returns a function that
// accepts another number (number2)
// 
// So sum2Numbers(2)(3) will return 5 and sum2Numbers(20)(10) will return 30
func sum2Numbers(number1 int) func (int) int {
  // Return another function that actually uses number1, you can 
  // see that this function here has access to the scope of sum2Numbers.
  // This is called a closure.
  return func (number2 int) int {
  	return number1 + number2
  }
}

func main() {
	sum3Numbers(1)(2)(3) // Returns 6
	sum3Numbers(5)(2)(5) // Returns 12
	sum3Numbers(-2)(2)(0) // Returns 0
	
	sum2Numbers(2)(4) // Returns 6
	sum2Numbers(5)(8) // Returns 13
	sum2Numbers(259)(-250) // Returns 9

	// You don't have to call them all at once
    foo := sum2Numbers(10)

    foo(2) // 12
    foo(50) // 60
    foo(20) // 30
}
```

These are silly examples, but they show
that you can construct powerful code contraptions with just functions.

For more information on the concepts discussed:

- [Higher order functions](https://www.golangprograms.com/higher-order-functions-in-golang.html)
- [Closures](https://tour.golang.org/moretypes/25)
- [General article + currying](https://codeburst.io/callbacks-closures-and-currying-3cc14300686a)

---

## 'Inheritance' -> Embedding

If 2 structs have common properties or methods, you can
use something called embedding.

```go
package main

type CommonEntity struct {
  ID        string
  CreatedAt int
  UpdatedAt int
}

// All properties of CommonEntity are shared with Person
type Person struct {
  CommonEntity
  Name string
}

// All properties of CommonEntity are shared with Server
type Server struct {
  CommonEntity
  Model string
}
```

It's useful for mocking in case you only want to mock one method.

```go
package main

type RealInterface interface {
  Method1()
  Method2()
  Method3()
  Method4()
  Method5()
  Method6()
  Method7()
}

type RealObject struct {
  // Pretend that there are 7 struct methods here
}

// In your test

type MockObject struct {
  // Inherit all 'real' methods
  RealObject
}

// We only override this method, we don't have to override the other 6
func (m *MockObject) Method1() {
  // Do some mocking here
}

```

[Click here to find out more on type embedding](https://travix.io/type-embedding-in-go-ba40dd4264df)

---

## Dockerfiles

This workshop has nothing to do with Docker, but we did want to highlight
a very cool thing you can do with Go binaries and Docker.

Ever heard of `FROM scratch` in Docker?
It's the most minimal base image available, it can't be run or pulled, and 
it contains 0 filesystem layers.
It has no shell or utilities and a very low attack surface.

Go can statically compile programs to a single binary, allowing
you to copy-paste the binary into a `scratch` image and run it
without additional runtime environments.

Below is an example dockerfile.

```dockerfile
##########################
# STEP 1 build the program
##########################
FROM golang:1.14-alpine AS builder

# Git is required for fetching the dependencies.
RUN apk update \
 && apk add --no-cache git

# See https://stackoverflow.com/a/55757473/12429735RUN
# We create a user so that we don't run our container as root
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "100001" \
    "appuser"

# Make sure everything is copied here
WORKDIR $GOPATH/src/package

# Copy all our files
COPY . .

# Fetch dependencies.
RUN go get -u -d -v

# Go build our binary, specifying the final OS and architecture we're looking for
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/my-program

############################
# STEP 2 build a small image
############################

# Here it is :-)
FROM scratch

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable.
COPY --from=builder /go/bin/my-program /go/bin/my-program

# Use the user that we've just created, one that isn't root
USER appuser:appuser

# Run the my-program binary.
ENTRYPOINT ["/go/bin/my-program"]
```
