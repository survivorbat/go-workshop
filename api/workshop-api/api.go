package workshop_api

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// Return a list of people
func getPeople() ([]string, error) {
	return nil, errors.New("not implemented")
}

// This route returns a list of people from the 'basics' file
func GetPeopleRoute(c *gin.Context) {
	// Ignore the error
	result, _ := getPeople()

	// Return json
	c.JSON(200, result)
}
