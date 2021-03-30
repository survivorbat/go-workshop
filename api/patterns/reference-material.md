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


## Logging

In the Basics part we used `fmt.Println()` to print messages to the
console, this is useful for make-shift debugging but it is not
actual logging.

To make logging possible, it's useful to use a logger.
Luckily Go has a built-in logger for this.
[Click here to learn more about this package](https://golang.org/pkg/log/).

There are also more packages available, [click here for examples](https://www.datadoghq.com/blog/go-logging/).

## Higher order functions

// TODO

## 'Inheritance' -> Embedding

If 2 structs have common properties or methods, you can
use something called embedding.

```go
type CommonEntity struct {
  ID        string
  CreatedAt int
  UpdatedAt int
}

// All properties of CommonEntity are shared with Person
type Person struct {
  CommonEntity
  // ... More properties
}

// All properties of CommonEntity are shared with Database
type Database struct {
  CommonEntity
  // ... More properties
}
```

It's useful for mocking in case you only want to mock one method.

```go
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
