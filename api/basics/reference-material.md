# The Basics

The following topics will be discussed in the 'basics' portion of the
workshop.

Please note that these are all simple examples and the language offers
many more useful tips and tricks.
Links have been added to each section for more information and examples.

- [The Basics](#the-basics)
	- [Hello World](#hello-world)
	- [Packages](#packages)
		- [Dependencies](#dependencies)
	- [Types](#types)
	- [Variable declarations](#variable-declarations)
		- [Variable casing](#variable-casing)
	- [Constants](#constants)
	- [If, for and switch](#if-for-and-switch)
		- [Example if statement](#example-if-statement)
		- [Example for loop](#example-for-loop)
		- [Example switch statement](#example-switch-statement)
	- [Arrays](#arrays)
	- [Maps](#maps)
	- [Functions](#functions)
	- [Structs](#structs)
	- [Pointers](#pointers)
	- [Interfaces](#interfaces)
	- [The null of Go](#the-null-of-go)
	- [Errors](#errors)
	- [Testing](#testing)

## Hello World

Before we get to the basics, we want to show you what a 'Hello World'
application looks like in Go.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

The entry point of an application is the `main` function in the `main` package.

---

## Packages

Every go file starts with a package declaration.

```go
package api
```

Just like other languages, Go codebases are divided into packages. File `basics/route.go` can access all objects
from `basics/basics.go` and vice-versa.
They are both in the `basics` package.

The `/main.go` file however, is part of the `main` package and can not directly access these objects and needs
to import `api` as a package.

```go
package main

import (
	// This package is housed in github.com/survivorbat/go-workshop.git so that's what
	// we're importing it from
	"github.com/survivorbat/go-workshop.git/api/basics"
)
```

- [Click here for more information on Go packages](https://www.callicoder.com/golang-packages/)
- [Click here for more information on modules](https://blog.golang.org/using-go-modules)

---

### Dependencies

Dependencies can also imported be using URLs, which makes it very
easy to download other packages.

```go
package main

import (
	"github.com/gin-gonic/gin"
	
	// Adding aliases is also possible
	logging "github.com/op/go-logging"
)
```

We're not diving further into this topic in this workshop.

[Click here for more information on modules](https://blog.golang.org/using-go-modules)

---

## Types

Go is a strongly typed language.
In this workshop we'll mostly use ints and strings, but there are plenty more to chooes from.

For more information:

- [Types](https://tour.golang.org/basics/11)
- [Zero values](https://tour.golang.org/basics/12)
- [Type conversions](https://tour.golang.org/basics/13)
- [Type inference](https://tour.golang.org/basics/14)

---

## Variable declarations

There are multiple ways to declare variables in Go.

```go
package main

// Using the 'var' keyword, allows you to explicitly define the type
var name1 string
var name2 string = ""
var name3 = ""

// Var can also be used in bulk, allowing you to specify multiple variables
// at the top of a file
var (
	name4 string
	name5 string = ""
	name6 = ""
)

// You can define multiple variables at once
var a, b, c = "a", "b", "c"

func myFunction() {
	// Shorthand variable syntax, only works in functions
	name7 := ""
}
```

[Click here for more info on variables](https://gobyexample.com/variables)

---

### Variable casing

Variable, property and function casing has great significance in Go.
Variables and properties starting with a capital letter are exported in the package.

```go
package main

// This variable can be used by other packages
var Exported = "This variable is exported"

// This variable is for this package only
var notExported = "This variable is not exported"
```

[Click here for more info on naming conventions](https://betterprogramming.pub/naming-conventions-in-go-short-but-descriptive-1fa7c6d2f32a)

---

## Constants

Constant values with a syntax similar to `var`.
Useful for avoiding magic numbers or strings and declaring addresses or
URLs in your code.

```go
package main

const name1 string = ""
const name2 = ""

const (
	name3 string = ""
	name4 = ""
)
```

For more information:

- [Constants](https://gobyexample.com/constants)
- [Numeric Constants](https://tour.golang.org/basics/16)

---

## If, for and switch

### Example if statement

```go
package main

import "fmt"

func main() {
	title := "Operations"

	// If statements have no parentheses
	if title == "Operations" {
		fmt.Println("I am an Operations")
		
	} else if title == "Product Owner" {
		fmt.Println("I am a Product Owner")
		
	} else {
		fmt.Println("I am NOT an Operations")
	}
}
```

### Example for loop

```go
package main

import "fmt"

func main() {
	// An example for loop with numbers, there are many variations and options
	for index := 1; index < 5; index++ {
		fmt.Printf("Number is: %v", index)
	}
}
```

`range` allows you to easily write for-loops by doing the hard-work
of iterating through certain types.
Take the following example.

```go
package example

func main() {
   numbers := []int{1, 2, 3}
   
   // Turn this into a for-each loop with range, you can
   // omit 'index' by making it an underscore (_)
   for index, number := range numbers {
	   	// Print the results
	   	fmt.Println(index)
		fmt.Println(number)
   }
}
```

### Example switch statement

```go
package main

import "fmt"

func main() {
	title := "Operations"

	// Switch statements look like this
	switch title {
	case "Operations":
		fmt.Println("I am an Operations")

	case "Developer":
		fmt.Println("I am a Developer")

	default:
		fmt.Println("I don't know what I am")
	}
}
```

- [Click here for some more examples on if statements](https://gobyexample.com/if-else)
- [Click here for some more examples on for loops](https://yourbasic.org/golang/for-loop/)
- [Click here for some more exmaples on `range`](https://gobyexample.com/range)
- [Click here for some more examples on switch statements](https://gobyexample.com/switch)

---

## Arrays

Arrays in Go are declared using square brackets.

```go
package main

// A simple string array
var myStrings = []string{"foo", "bar"}

// A multi-dimensional integers
var myMultiInts = [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

// An array of Person objects
type Person struct {
    Name string
    Title string
}

var myPeople = []Person{
	{Name: "Janneth", Title: "Product Owner"},
	{Name: "Rob", Title: "Developer"},
}
```

[Click here for more examples on arrays](https://gobyexample.com/arrays)

[Click here to read more about slices](https://gobyexample.com/slices)

---

## Maps

Go has maps and are declared as: `map[<key-type>]<value-type>`.

```go
package main

// A simple string-to-string map
var myStrings = map[string]string {
	"hello": "world",
}

// A bool-to-int map
var myObjects = map[bool]int {
	true: 23,
	false: 102,
}

// A string-to-struct map where we create a new struct at
// creation of the map.
// More on structs later ;-)
var myStructs = map[string]struct{
	SomeValue string
}{
	{SomeValue: "Hello"},
	{SomeValue: "World"},
}
```

[Click here for more examples on maps](https://gobyexample.com/maps)

---

## Functions

In Go you write functions using the `func` keyword.

```go
package main

// A simple function
func myFunction() {
	// Do something
}

// A function with multiple input parameters
func myFunctionWithParameters(a string, b int, c bool, d map[string]string) {
	// Do something
}

// A function with a return value
func myFunctionWithReturnType() string {
	return "Hello World"
}

// A function with multiple return values
func myFunctionWithMultipleReturnTypes() (string, int, bool) {
	return "Hello World", 32, true
}

// Calling functions
func main() {
	// Calling functions
	myFunction()
	myFunctionWithParameters("", 20, true, map[string]string{
		"foo": "bar",
    })
	
	// Getting one return value
	helloWorldString := myFunctionWithReturnType()
	
	// Getting multiple return values
	myString, myInt, myBool := myFunctionWithMultipleReturnTypes()
	
	// Ignoring some return types can be done with underscores
	_, _, someBool := myFunctionWithMultipleReturnTypes()
}
```

You can find more examples here:

- [Functions](https://gobyexample.com/functions)
- [Multiple return values](https://gobyexample.com/multiple-return-values)
- [Variadic functions](https://gobyexample.com/variadic-functions)
- [Closures](https://gobyexample.com/closures)
- [Recursion](https://gobyexample.com/recursion)

---

## Structs

Go does not have classes but uses structs.
Structs can have both methods and properties.

```go
package main

import "fmt"

// A Person struct
type Person struct {
	// The name of a person
	Name string

	// Title of the person
	Title string
}

// Now, we add a method to the Person struct
func (p Person) Speak() {
	// Print the person's name to the screen
	fmt.Sprintln("Hello my name is %v", p.Name)
}

// Methods are functions with extra parenthesis at the front,
// just like functions they have parameters and return types
func (p Person) Say(message string) {
	// Print the message with the person's name
	fmt.Sprintln("%v says: '%v'", p.Name, message)
}
```

A struct can be instantiated using curly braces.

```go
package main

type Person struct {
    Name string
}

var myPerson Person = Person{}
```

You can find more examples here:

- [Structs](https://gobyexample.com/structs)
- [Methods](https://gobyexample.com/methods)

---

## Pointers

Pointers are useful if you want to pass data by reference instead of value.
Take the following example.

```go
package main

import "fmt"

// A person object
type Person struct {
	Name string
}

func main() {
	// Create Jennifer
	person := Person{
		Name: "Jennifer",
    }
    
    // Attempt to change their name to Bob
    changeNameToBob(person)
	
	// Result: Jennifer
	fmt.Println(person.Name)
}

func changeNameToBob(person Person) {
	person.Name = "Bob"
}
```

This does not change the name of the person.
To get this working properly, you need to change the parameter type of `changeNameToBob`
to a pointer.

```go
package main

import "fmt"

// A person object
type Person struct {
	Name string
}

func main() {
	// Create Jennifer
	person := Person{
		Name: "Jennifer",
    }
    
    // Get the pointer to the Person object using the & operator
    personPointer := &person
    
    // Attempt to change their name to Bob
    changeNameToBob(personPointer)
	
	// Result: Bob
	fmt.Println(person.Name)
}

// The parameter here is a pointer type, signified by the star *
func changeNameToBob(person *Person) {
	person.Name = "Bob"
}
```

[Click here to find more examples on pointers](https://gobyexample.com/pointers)

---

## Interfaces

Go interfaces are similar to interfaces in other languages, they
define expected behaviour of objects.
One big difference is, that you do not explicitly define what interface
a struct belongs to.

```go
package main

type PersonInterface interface {
	Say(message string)
}

type Person struct {
    Name string	
}

func (p Person) Say(message string) {
	// Do something
}

// The person object adheres to the PersonInterface and allows us to use
// it like this. We can't declare that the interface belongs to a concrete type
// like in other languages.
var person PersonInterface = Person{Name: "Jerry"}
```

[Click here for more information on Interfaces](https://gobyexample.com/interfaces)

---

## The null of Go

Go uses `nil` as a null pointer, it's similar to `null` in other languages.

[Click here for more information on Nil](https://go101.org/article/nil.html)

---

## Errors

In other languages you might be using to exceptions to handle unexpected
application states that are caught in a method further down the stack.

Go on the other hand, uses explicit error return types.
This means that your functions will often return either an error or a return type
and an error.

```go
package main

import "errors"

func doThing(someString string) error {
	// If the string is equal to "please work", we're NOT returning an error
	if someString == "please work" {
		return nil
    }
    
    // Otherwise, we return an error based on a string
	return errors.New("something went wrong")
}

func validatePositiveNumber(number int) (int, error) {
	// If the number is lower or equal to 0, we're returning 0 and an error
	if number <= 0 {
		return 0, errors.New("number was not positive")
	}
	
	// Otherwise we return the number and no error
	return number, nil
}
```

There's also a `panic()` and a `recover()` function, `panic()` is for unexpected
errors that can NOT be handled grafefully.
For example when logging can not be initialized or when environment variables are missing at startup.

Always go with returning an error unless there is no other option.

[Click here for more information on errors in Go](https://gobyexample.com/errors)

[Click here for more information on panic in Go](https://gobyexample.com/panic)

## Testing

Tests in Go are written in `_test.go` files.
A test function starts with the word Test and has 1 parameter, `t *testing.T`.

```go
package main

// main.go

// Our example function that we want to test
func HelloWorld() string {
	return "Hello World"
}

// main_test.go

// Starts with test and has 1 parameter
func TestHelloWorld_ReturnsHelloWorld(t *testing.T) {
	// Act
	result := HelloWorld()

	// Assert
	assert.Equal(t, "Hello World", result)
}
```

[Click here for a full guide on writing tests](https://blog.alexellis.io/golang-writing-unit-tests/).

*Please note that this guide does not follow all the practices we use in the workshop*
