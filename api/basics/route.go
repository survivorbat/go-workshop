package basics

import (
	"github.com/gin-gonic/gin"
)

/*
	This file contains routes for the API that we're building.
	These functions all contain gin (web framework) specific code and call upon
	functions in the rest of the Go files.
*/

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
