# Patterns Exercises

Please feel free to stray from the path and do your thing,
part of the fun is discovering Go yourself :-)

Run `make r` or `gin --appPort 8080 -i` in the root directory to start the application.

Then you can visit these pages:

- [localhost:8080/api/people](http://localhost:8080/api/people)

- [Patterns Exercises](#patterns-exercises)
  - [Exercise 1, Mocking Routes](#exercise-1-mocking-routes)
    - [A: Controller struct](#a-controller-struct)
    - [B: Service struct](#b-service-struct)
    - [C: Interface magic](#c-interface-magic)
    - [D: Almost there](#d-almost-there)
    - [Intermission](#intermission)
    - [E: Writing a test](#e-writing-a-test)
    - [F: The POST route](#f-the-post-route)

## Exercise 1, Mocking Routes

Right now we're using plain functions to handle our routes,
these functions use other functions and variables to perform
their tasks.

Let's change that and make these functions mockable.

### A: Controller struct

Let's first create a struct.

```go
type PeopleController struct { 
}
```

Change the GetPeopleRoute so that it becomes
part of the struct.

```go
func (p *PeopleController) GetPeopleRoute(c *gin.Context) {
    
}
```

This will throw errors in the `main.go`, to fix these
we have to instantiate the PeopleController and add its function to the routes:

```go
peopleController := workshop_api.PeopleController{}

apiRunner.GET("/api/people", peopleController.GetPeopleRoute)
```

There we go, the controller is now a struct.
How does this help us testing it?

### B: Service struct

Right now the `getPeople` and `addPerson` are floating around in the `api.go` file,
let's organize them in a service.

Create a PeopleService struct:

```go
type PeopleService struct {}
```

Add the methods:

```go
func (p *PeopleService) getPeople(titleFilter string) ([]Person, error) {}
func (p *PeopleService) addPerson(person Person) error {}
```

If you like, you can put the People variable in the struct, but you don't have to.

### C: Interface magic

Now, create the following interface:

```go
type PeopleServiceInterface interface {
    getPeople(titleFilter string) ([]Person, error)
    addPerson(person Person) error
}
```

This will be the public API for this service, we expose these 2 methods
and every struct that handles people will have to adhere to these
method signatures.

In any other language you'd probably write `PeopleService implements PeopleServiceInterface`,
but Go does not have such a construct.

Now, change your PeopleController to something like this:

```go
type PeopleController struct { 
    PeopleService PeopleServiceInterface
}
```

And in your methods, instead of calling `getPeople` or `addPerson` directly,
you'll be using the following construct:

```go
func (p *PeopleController) GetPeopleRoute(c *gin.Context) {
    // [...]
    result, err := p.PeopleService.getPeople(titleFilter)
    // [...]
}

func (p *PeopleController) PostPeopleRoute(c *gin.Context) {
    // [...]
    err := p.PeopleService.addPerson(postedPerson)
    // [...]
}
```

### D: Almost there

The last step is to go to the `main.go` and change the following code.

```go
peopleController := workshop_api.PeopleController{
    // Instantiate a people service
    PeopleService: &workshop_api.PeopleService{}
}

apiRunner.GET("/api/people", peopleController.GetPeopleRoute)
```

Now you should be able to run the application again and have everything working.

### Intermission

Okay so what did we achieve now? These changes make the code more complicated,
what's the benefit? Decoupling.

Our routes are no longer reliant on the `getPeople` or `addPerson` functions,
they're reliant on an interface.
This is called [Dependency Inversion](https://www.tutorialsteacher.com/ioc/dependency-inversion-principle).

We can give the `PeopleController` any object as long as it implements those 2 methods.
A `PeopleDatabaseService`, a `PeopleApiService` or even a `PeopleExcelService`, as long
as it's compatible with the interface it doesn't matter.

But also, a `MockPeopleService`.
In the next section, we're finally going to write a proper test for this controller.

### E: Writing a test

Let's create the `MockPeopleService` in `api_test.go`:

```go
type MockPeopleService struct {
    GetPeopleReturns []Person
    GetPeopleError error
    GetPeopleCalledWith string

	AddPersonError error
	AddPersonCalledWith Person
}

func (p *MockPeopleService) getPeople(filterTitle string) ([]Person, error) {
	p.GetPeopleCalledWith = filterTitle

	return p.GetPeopleReturns, p.GetPeopleError
}

func (p *MockPeopleService) addPerson(person Person) error {
	p.AddPersonCalledWith = person
	return p.AddPersonError
}
```

This service will allow us to properly test whether
the routes behave accordingly whenever edge-case values are returned.

For example:

```go
func TestGetPeopleRoute_Returns500OnError(t *testing.T) {
	tests := map[string]struct{
		errorMessage string
	} {
		"test": {errorMessage: "test"},
		"error occurred": {errorMessage: "error occurred"},
	}

	for name, testData := range tests {
		t.Run(name, func (t *testing.T) {
			// Arrange

			// Make sure this returns an error
			mockService := MockPeopleService{}
			mockService.GetPeopleError = errors.New(testData.errorMessage)

			// Get the controller
			controller := PeopleController{PeopleService: &mockService}

			// Response will be written to this writer
			writer := httptest.NewRecorder()

			// Test context for Gin
			c, _ := gin.CreateTestContext(writer)

			// Add a request to the context, so we can extract values from it (title)
            c.Request = httptest.NewRequest(http.MethodGet, "https://example.com?title=empty", nil) 

            // Act
			controller.GetPeopleRoute(c)

			// Assert
			assert.Equal(t, 500, writer.Code)
		})
	}
}
```

Now, create a test like this for the GetPeopleCalledWith property,
check if the function is called with the expected parameter (title).

