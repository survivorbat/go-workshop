package workshop_api

import "github.com/gin-gonic/gin"

/**
Exercise 1, 2, 3 and 4
*/

type Person struct {
	Name string `json:"name,omitempty"`
	Title string `json:"title,omitempty"`
}

var People = []Person{
	{Name: "Lindsay", Title: "Developer"},
	{Name: "Bob", Title: "Product Owner"},
	{Name: "Chris", Title: "Developer"},
	{Name: "Dahlia", Title: "Operations"},
}

//// Return a list of people
//func getPeople(title string) ([]Person, error) {
//	var result []Person
//
//	for _, person := range People {
//		if person.Title == title {
//			result = append(result, person)
//		}
//	}
//
//	return result, nil
//}
//
//// This route returns a list of people from the 'basics' file
//func GetPeopleRoute(c *gin.Context) {
//	titleFilter := c.Query("title")
//
//	// Ignore the error
//	result, err := getPeople(titleFilter)
//
//	if err != nil {
//		c.JSON(500, "an error occurred!")
//		return
//	}
//
//	// Return json
//	c.JSON(200, result)
//}
//
//func addPerson(person Person) error {
//	People = append(People, person)
//	return nil
//}
//
//func PostPeopleRoute(c *gin.Context) {
//	var postedPerson Person
//
//	// Note here that we're declaring a variable within an if-statement,
//	// this is a valid Go construct :-)!
//	if err := c.ShouldBindJSON(&postedPerson); err != nil {
//		c.JSON(400, map[string]string{ "error": err.Error() })
//		return
//	}
//
//	err := addPerson(postedPerson)
//
//	if err != nil {
//		c.JSON(500, map[string]string{"error": err.Error()})
//		return
//	}
//
//	c.JSON(200, map[string]string{"result": "Success!"})
//}

/**
Patterns exercise 1
 */

type PeopleController struct {
	PeopleService PeopleServiceInterface
}

// This route returns a list of people from the 'basics' file
func (p *PeopleController) GetPeopleRoute(c *gin.Context) {
	titleFilter := c.Query("title")

	// Ignore the error
	result, err := p.PeopleService.getPeople(titleFilter)

	if err != nil {
		c.JSON(500, "an error occurred!")
		return
	}

	// Return json
	c.JSON(200, result)
}

func (p *PeopleController) PostPeopleRoute(c *gin.Context) {
	var postedPerson Person

	// Note here that we're declaring a variable within an if-statement,
	// this is a valid Go construct :-)!
	if err := c.ShouldBindJSON(&postedPerson); err != nil {
		c.JSON(400, map[string]string{ "error": err.Error() })
		return
	}

	err := p.PeopleService.addPerson(postedPerson)

	if err != nil {
		c.JSON(500, map[string]string{"error": err.Error()})
		return
	}

	c.JSON(200, map[string]string{"result": "Success!"})
}

type PeopleServiceInterface interface {
	getPeople(titleFilter string) ([]Person, error)
	addPerson(person Person) error
}

type PeopleService struct {}

// Return a list of people
func (p *PeopleService) getPeople(title string) ([]Person, error) {
	var result []Person

	for _, person := range People {
		if person.Title == title {
			result = append(result, person)
		}
	}

	return result, nil
}

func (p *PeopleService) addPerson(person Person) error {
	People = append(People, person)
	return nil
}
