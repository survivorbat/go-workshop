# API

The following topics will be discussed in the 'api' portion of the
workshop.

Please note that these are all simple examples and the language/framework offers
many more useful tips and tricks.

- [API](#api)
  - [API Framework](#api-framework)
  - [Main.go](#maingo)
  - [Gin Context](#gin-context)
  - [Unserializing JSON](#unserializing-json)

---

## API Framework

In the previous section of the workshop we showed you how to serialize JSON and return
it in a function.
We promised you that it'd be come easier, so let's get right to the juice.

The Go community offers you [a lot of options](https://medium.com/devtechtoday/top-7-golang-web-frameworks-in-2020-and-beyond-9ca2a89eb904)
when it comes to web frameworks.
Some - [like Gorilla Mux](https://github.com/gorilla/mux) - provide you with bare-bone utilities
as to allow you to customize your application to the fullest while others - [like Gin](https://github.com/gin-gonic/gin)
- offer you a full-fledged framework.

To not have to write and maintain utility code ourselves, we opted to use the Gin framework for
the workshop.
The Gin documentation is pretty extensive and covers most of the use-cases that we're going to use,
still, we wanted to give you some useful blocks of code to get going.

---

## Main.go

The `main.go` file at the root has all the code needed to set up
the API of this workshop.

It first instantiates a default Gin instance (`gin.Default()`) which has
built-in logging and error handling.
Next the routes are added to the instance, in our case our People routes and
the the basics page.
At last, we run the server on an address and log any errors.

The .GET, .POST, .DELETE, .PATCH etc. routes all take 2 arguments, a path and
a function.
This function has to have 1 input parameter which is a pointer to a gin.Context
object like so: `myRoute(c *gin.Context)`.

Take a step back and go to the `basics` folder, there's a `route.go` file that was
not mentioned in the course material.
It contains the function that rendered the HTML page you used in the exercises.

---

## Gin Context

The *gin.Context object is extremely useful and takes way
a lot of manual work like serializing and reading parameters
from a URL.

For example, to send responses.

```go
func myRoute(c *gin.Context) {
    // HTML
    c.HTML(200, "index.html")
    return

    // JSON
    c.JSON(200, {})
    return

    // String
    c.String(200, "hello world")
    return
}
```

Or to get parameters from the query.

```go
func myRoute(c *gin.Context) {
    someInput := c.Query("someInput")
}
```

---

## Unserializing JSON

Gin also offers you a neat way of unserializing JSON into an object.

```go
func myRoute(c *gin.Context) {
    var postedPerson Person

    // Note here that we're declaring a variable within an if-statement,
    // this is a valid Go construct :-)!
    if err := c.ShouldBindJSON(&requestObject); err != nil {
        // Do something if the input data was faulty
    }

    // Do something with the person
}
``` 

Gin even has Struct tags that allow you to validate the incoming data, pretty cool.
