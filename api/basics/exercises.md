# The Basics Exercises 

Please feel free to stray from the path and do your thing,
part of the fun is discovering Go yourself :-)

Run `make r` or `gin --appPort 8080 -i` in the root directory to start the application.

Then you can visit these pages:

- [localhost:8080/basics/foo](http://localhost:8080/basics/foo)
- [localhost:8080/basics/bar](http://localhost:8080/basics/bar)
- [localhost:8080/basics/baz](http://localhost:8080/basics/baz)

The final parameter (foo/bar/baz) corresponds to the `name` parameter
in the `getResult()` function in `basics.go`.
You can change this to anything you like.

- [The Basics Exercises](#the-basics-exercises)
  - [Before you begin](#before-you-begin)
  - [Exercise 1, string formatting](#exercise-1-string-formatting)
    - [A: Hello World](#a-hello-world)
    - [B: Hello name](#b-hello-name)
    - [C: Printing](#c-printing)
    - [D: Length](#d-length)
    - [E: Errors](#e-errors)
  - [Exercise 2, conditionals](#exercise-2-conditionals)
    - [A: If](#a-if)
    - [B: Else, if](#b-else-if)
    - [C: Switch](#c-switch)
  - [Exercise 3, Loops and Functions](#exercise-3-loops-and-functions)
  - [Exercise 4, JSON and structs](#exercise-4-json-and-structs)
    - [A: Getting a variable](#a-getting-a-variable)
    - [B: Map](#b-map)
    - [C: Mapping results](#c-mapping-results)
    - [D: Serializing](#d-serializing)
    - [E: Structs](#e-structs)
    - [F: Tags](#f-tags)
  - [Exercise 5, Writing a test](#exercise-5-writing-a-test)
    - [A: Test function](#a-test-function)
    - [B: Test data](#b-test-data)
    - [C: Looping](#c-looping)
    - [D: Writing the method](#d-writing-the-method)
    - [E: Error test](#e-error-test)
  - [Conclusion](#conclusion)

## Before you begin

Go has a few compile-check in place that force you to
write better code.
One of these checks is for unused variables.
If you declare a variable but don't use it, Go will call you
out on it and refuse to compile.

If a function returns a variable you don't want to use, you can use
underscores (_).

## Exercise 1, string formatting

### A: Hello World

Instead of an empty string, make the 'getResult' function print 'Hello World' to the screen.

### B: Hello name

Define a variable called 'resultString' and set it equal to fmt.Sprintf("Hello %v", name),
return this variable in getResult and see what it outputs to the screen

Change the 'foo' in the URL to something else and see what happens.

The fmt package is very useful for printing to the screen and formatting strings,
find more info here: https://golang.org/pkg/fmt/

### C: Printing

Let's add some make-shift logging, add fmt.Println("") to your getResult function with a fun message.
Reload the page and see what shows up in your logs (where you ran `make r`).

Note that this isn't _real_ logging, but it demonstrates writing to a console.

### D: Length

Now we want to return something like: "Input string 'Foo' has '3' characters."
You can get the length of a string using the `len()` function.

Re-use the fmt.Sprintf() method with more inputs to create the string you need to return.

### E: Errors

Now let's try returning an error, change the `nil` to
errors.New("") with a nice error string.

## Exercise 2, conditionals

### A: If

We've been returning simple strings now, let's add some actual logic.

Change the code so that if the input string equals `world`, it returns `hello`.
If the input equals anything other than `world`, it should return
an error with some nice text.

- [localhost:8080/basics/world](http://localhost:8080/basics/world) should return 'hello'
- [localhost:8080/basics/something-else](http://localhost:8080/basics/something-else) should return an error

### B: Else, if

Add another condition that checks if the input equals `hello`, if so it should return 'world'.

Use the `} else if ... {` construct for this.

- [localhost:8080/basics/world](http://localhost:8080/basics/world) should return 'hello'
- [localhost:8080/basics/hello](http://localhost:8080/basics/world) should return 'world'
- [localhost:8080/basics/something-else](http://localhost:8080/basics/something-else) should return an error

### C: Switch

Using if-else statements is an option, but it might be better to use a switch statement for this.
Set the `default` case to return an error.

## Exercise 3, Loops and Functions

Now for something a bit more interesting, let's start reversing the input string.

We want to have the following results:

| Input | Result
| ----- | ------
| Hello | olleH
| Good morning! | !gninrom dooG
| Ok | kO
| t | Error
| b | Error

Let's put this functionality in another function.
Write a function below the `getResult()` and call it something similar to `reverseString`, it should
have one string input parameter and return either a string and a potential error.

First, let the `reverseString()` function return an error if the input string is shorter than 2 characters,
there's nothing to reverse if the string is too short.
An if-statement is a good candidate for this.

Now initialize a result string and write a for-loop that uses `range` on the input string.
Loop through the input string and add every character you find to the result string in opposite order,
then return the result.

In your `getResult`, write `return reverseString(...)` so that the 2 return values of `reverseString` are
returned as the function result.

## Exercise 4, JSON and structs

In the next section of the workshop, we'll be using a web framework for JSON serialization.
Since we're not there yet we'll be doing it manually for the sake of learning.

We'll be continuing with the previous exercise, but instead of just returning
the result, we want to have something like:

| Input | Result
| ----- | ------
| Hello | `{ "result": "olleH" }`
| Good morning! | `{ "result": "!gninrom dooG" }` 
| Ok | `{ "result": "kO"}`
| t | `{ "error": "Your input has to be at least 2 characters" }`
| b | `{ "error": "Your input has to be at least 2 characters" }`

Luckily, we don't have to do this from scratch completely, we have the
`json` package for that.

We're going to use the `func Marshal(v interface{}) ([]byte, error)` function that you can use without having to download anything.

As you can see in the signature, the input is an `interface{}` which is
a wildcard for _any_ type.

### A: Getting a variable

Capture the `reverseString()` result in a variable: `stringResult, reverseError := reverseString(name)`.

### B: Map

Create a map variable: `resultMap := map[string]string{}`

### C: Mapping results

Now make an if-statement that checks whether `reverseError` equals nil, if it does, set `resultMap["result"] = stringResult`.
Otherwise, `resultMap["error"] = err.Error()` which will set ["error"] to the error string.

### D: Serializing

Now that we have our map, convert it to json by using `json.Marshal(resultMap)`.
Capture the result and discard the error with `jsonResult, _ := ...`.

Marshal returns bytes, you can cast it back to a string by using `string()` around
the result from `json.Marshal`.
Now return it from the `getResult` function.

Try a few inputs and see if the errors appears in 'error' and the result in 'result'.

**For this exercise you no longer need to return an error from getResult, the error is now in the JSON result object**

### E: Structs

So now we used a map to gather our values, let's be a bit more _struct_-ured and start
putting the data in a struct.
Above the `getResult` function, define a struct using `type Response struct {}`. 

Give it 2 properties, `Result` and `Error`, make sure to capitalize
the first character. 
If you don't, the properties won't be serialized.
They should both be `string` types.

Replace the mapResult with `Response{}` and the assignments with
the properties you created.

Now the result should look something like this: `{ "Result": "!olleH", "Error": null }`

Almost correct, Result is capitalized and "Error" is there.

### F: Tags

To fix this, we can add tags to our struct. Make the following
changes to your struct:

```go
package basics

type Response struct {
    Result string `json:"result,omitempty"`
    Error string `json:"error,omitempty"`
}
```

This will omit empty values and make sure the casing of the
result is correct.

There we are! A nice struct result.

[Click here for more information on tags](https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go)

## Exercise 5, Writing a test

So far we've been naughty, we haven't written a single unit test for our code!
As a result, we can not be sure that everything works as expected.
Let's open `basics_test.go` and write a unit test for `getResult()`.
In this phase we'll be using 2 principles that help us writing effective
unit-tests: [The triple-A pattern](https://wiki.c2.com/?ArrangeActAssert) and data-driven tests.

Use `make t` to run the unit tests.

### A: Test function

First, create the function that will be our test.
The signature will be `Test_GetResult_ReturnsExpectedResult(t *testing.T)`.

A test function always starts with the word 'Test' and has one parameter,
a pointer to a *testing.T instance.

### B: Test data

Now we're going to write some test data using a map and a struct.

```go
package basics

import "testing"

// The test method
func Test_GetResult_ReturnsExpectedResult(t *testing.T) {
    // Our test data, we provide both a name (the key) and
    // a struct that we create on the spot. It doesn't matter
    // what you call the key (hello, good, tacocat) as long as
    // they are unique.
    // This struct has 2 properties:
    //
    // input -> What we're going to give the getResult function
    // expectedOutput -> What we expect to get in return
    // 
    tests := map[string]struct{
        input string
        expectedResult string
    }{
        // First test data set, hello should return olleh
        "hello": {input: "hello", expectedResult: "olleh"},

        // Second test data set, good should return doog
        "good": {input: "good", expectedResult: "doog"},

        // Third test data set, a palindrome
        "tacocat": {input: "tacocat", expectedResult: "tacocat"},
    }
}
```

Good, now we know what our expectations are for the function
we've created.

### C: Looping

Let's use this data, first we need to loop through the
test data sets and run individual tests. To accomplish this, take
the following construct.

```go
package basics

// The test method
func Test_GetResult_ReturnsExpectedResult(t *testing.T) {
    tests := map[string]struct{
        input string
        expectedResult string
    }{
        "hello": {input: "hello", expectedResult: "olleh"},
        "good": {input: "good", expectedResult: "doog"},
        "tacocat": {input: "tacocat", expectedResult: "tacocat"},
    }

    // Use range to loop through the key-values, name is the key
    // and testData is the struct
    for name, testData := range tests {
        // Run the test using the name
        t.Run(name, func (t *testing.T) {
            
        })
    }
}
```

Great, we're almost there!

### D: Writing the method

Now we're going to write the actual test, what do we want to do?

1. Give the `input` to `getResult` and store the returning value and error
2. Check whether the error is `nil`
3. Check whether the result is equal to the expected result

Alright, let's finish this test:

```go
package basics

// The test method
func Test_GetResult_ReturnsExpectedResult(t *testing.T) {
    tests := map[string]struct{
        input string
        expectedResult string
	}{
		// First test data set, hello should return olleh
		"hello": {input: "hello", expectedResult: "{\"result\":\"olleh\"}"},

		// Second test data set, good should return doog
		"good": {input: "good", expectedResult: "{\"result\":\"doog\"}"},

		// Third test data set, a palindrome
		"tacocat": {input: "tacocat", expectedResult: "{\"result\":\"tacocat\"}"},
	}

    for name, testData := range tests {
        t.Run(name, func (t *testing.T) {
			// Act
			result, err := getResult(testData.input)

			// Assert
			if assert.Nil(t, err) {
				assert.Equal(t, testData.expectedResult, result)
			}
        })
    }
}
```

Notice here:

- We clearly note down what we're testing (ACT) and where in the code we are making assertions (ASSERT)
- Before we assert the result, we check if no error occurred to prevent nil errors

### E: Error test

Now it's your turn, write a test that asserts that an error
is returned if the input is only 1 character long.
Use a signature like 'ReturnsErrorOn1CharacterInput' or something
in that sense.

Make the following assertions:

- Assert that the error is nil
- Assert that the Error property of your struct contains the error message

Copy-paste the existing test as a head-start.

## Conclusion

Okay so we've seen a little of what the Go language has to offer,
you've used a few simple constructs and returned data to an already existing
HTML function.

In the next section, we'll actually start working on a functioning API.
