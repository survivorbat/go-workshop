# API Exercises

Please feel free to stray from the path and do your thing,
part of the fun is discovering Go yourself :-)

Run `make r` or `gin --appPort 8080 -i` in the root directory to start the application.

Then you can visit these pages:

- [localhost:8080/api/people](http://localhost:8080/api/people)

- [API Exercises](#api-exercises)
	- [Exercise 1, Window Shopping](#exercise-1-window-shopping)
	- [Exercise 2, Building GetPeopleRoute](#exercise-2-building-getpeopleroute)
		- [A: Handling an error](#a-handling-an-error)
		- [B: Creating people](#b-creating-people)
		- [C: Returning people](#c-returning-people)
		- [D: Testing `getPeople`](#d-testing-getpeople)
		- [E: Testing `getPeopleRoute`](#e-testing-getpeopleroute)
	- [Exercise 3: Filtering people](#exercise-3-filtering-people)
		- [A: Changing signatures](#a-changing-signatures)
		- [B: Update `getPeople` test](#b-update-getpeople-test)
		- [C: Test to see if it works](#c-test-to-see-if-it-works)
	- [Exercise 4: Creating people](#exercise-4-creating-people)
		- [A: Add the route](#a-add-the-route)
		- [B: Adding a person](#b-adding-a-person)
		- [C: Accept incoming POST requests](#c-accept-incoming-post-requests)
		- [D: Error handling](#d-error-handling)
		- [E: Adding people](#e-adding-people)
		- [F: Final touch](#f-final-touch)
		- [G: Check it out](#g-check-it-out)
	- [Conclusion](#conclusion)

## Exercise 1, Window Shopping

Before we start writing code, let's have a look at the [`route.go` file](../basics/route.go) from
the Basics portion.

Here, your `getResult` function was used to display the contents on the web page that
we provided.
Let's examine that code.

```go
package basics

func Basics(c *gin.Context) {
	// Get input
	input := c.Param("input")

	// Get result
	result, err := getResult(input)

	// Return the error as a 500 if anything goes wrong
	if err != nil {
		c.HTML(500, "basics.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return string result
	c.HTML(200, "basics.html", gin.H{
		"result": result,
	})
}
```

The function signature is very simple, it only required a pointer to a `gin` context.
This context is used several times to return HTML code and to get inputs from the request.

Returning HTML is pretty cool but quite cumbersome for a HTTP client to digest, luckily
there's also a `c.JSON` function.

## Exercise 2, Building GetPeopleRoute

Have a look at `api.go`'s `GetPeopleRoute` function, we've written a small block of code that
returns a list of people.
At least, it will in the future, now it returns `null`.

As you can already see, we've split the request/response handler (`GetPeopleRoute`) from
the actual logic that retrieves the people in `getPeople`.

This allows us to - in the future - test `GetPeopleRoute` separately from `getPeople`.

### A: Handling an error

Add an if statement that checks whether the second argument of `getPeople` returns an error,
if so, it should return a 500 error with a message.
Make sure to use `return` in the if statement to make sure the execution ends.

Reload your browser, it should now output the error message that you've chosen.

It's not JSON I know, but that'll come later.

### B: Creating people

Create a struct called `Person` and give it the properties `Name` and `Title`.
Don't forget to also add tags so that the JSON is serialized properly.

Now, create a variable of an array of people with different names and titles.

| Name | Title
| ---- | ----
| Lindsay | Developer
| Bob | Product Owner
| Chris | Developer
| Dahlia | Operations

Use something like:

```go
var People = []Person{
    {Name: "...", Title: "..." }
}
```

### C: Returning people

Change the `getPeople` function's signature and b ody to return the people you just created.
Reload your page, you should now see a list of people in JSON.

### D: Testing `getPeople`

Check out `api_test.go`.
Now that we've changed the result of `getPeople`, we have to change it.

Make sure that `err` is nil and the result equals the list of people you've created.

### E: Testing `getPeopleRoute`

Check out `api_test.go`

Now that we've changed the result of `getPeopleRoute`, we have to change it.

We can currently only test whether the route returns the expected JSON data,
we'll get into writing the error (500) test later.

Add the following assertions:

- Assert that `writer.Code` equals 200
- Assert that `writer.Body.string()` is equal to the JSON you expect (people)

## Exercise 3: Filtering people

Let's add filter parameters so that
we can filter the list of people based on a URL param.
Add the following to `GetPeopleRoute`:

```go
titleFilter := c.Query("title")
```

This will allow you to get `helloworld` if you go to [/api/people?title=helloworld](http://localhost:8080/api/people?title=helloworld).

### A: Changing signatures

Change `getPeople` so that the function filters the list of people based on the incoming `title` parameter.
If I want only Developers, `?title=Developer` should give me a list of only Developers.

Use your knowledge from the Basics part to finish this exercise.

### B: Update `getPeople` test

Now add testData to your getPeople test and check if you get the expected data
based on the supplied `title`.

Use the following construct:

```go
testData := map[string]struct{
	inputTile string
	expectedPeople []string
}{
	// Use your own data here ;-)
	{
		inputTitle: "Developer",
		expectedPeople: []string{"Dahlia", "Rob"},
	},
	{
		inputTitle: "Product Owner",
		expectedPeople: []string{"Bianca"},
	}
}
```

We'd also like you to update the GetPeopleRoute so that it
now also filters, but we'd like to ask you to wait until
we get to the next part of the workshop to make it easier.

### C: Test to see if it works

Open your browser and see if everything works as expected.
Hopefully everything's working nicely :-)

## Exercise 4: Creating people

OK so we're done with the GET route for now, let's start POST-ing.

You can use any tool you like to send a POST request to `/api/people`,
but in this workshop we're going to stick to simply using CURL.

### A: Add the route

Create a POST route like this:

```go
func PostPeopleRoute(c *gin.Context) {}
``` 

Now go to `main.go` at the root and add:

```go
apiRunner.POST("/api/people", workshop_api.PostPeopleRoute)
``` 

Now for testing, add `c.JSON(200, "success!")` to your function.
Now, run the shell script `post-person.sh` which will perform
a POST request to the route you just created.

Run: `./api/workshop-api/post-person.sh`

You should now see `success!`.

### B: Adding a person

Just like `getPeople`, now create a `addPerson(person Person) error` function that adds
the given person to the `People` list.
You can use the `append()` function for this: `People = append(People, person)`.

You may write a test for it as well: `TestAddPerson_AddsThePersonToTheList`

### C: Accept incoming POST requests

Almost there, remove the 'success' message.
Add the following piece of code to your `PostPeopleRoute`:

```go
var postedPerson Person

// Note here that we're declaring a variable within an if-statement,
// this is a valid Go construct :-)!
if err := c.ShouldBindJSON(&requestObject); err != nil {
	// TODO: Do something if the input data was faulty
}
```

The `postedPerson` variable will contain the person from the POST request.

### D: Error handling

Now add error handling at the comment, return
a 400 JSON error for example.
Make sure to put a `return` statement in to stop the function.

### E: Adding people

Now add a call to your `addPerson` function, remember to get the error back
from the function and put it in an `err` variable.

### F: Final touch

Check if `err != nil`.
It should return a 500 error or a 200 error depending on the error check.

### G: Check it out

Now, run the `post-person.sh` script again, does the person get added?

## Conclusion

Great, you just created a simple API.
The final part of the workshop will implement some useful patterns to

