package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/survivorbat/go-workshop.git/api/basics"
	"github.com/survivorbat/go-workshop.git/api/workshop-api"
)

const serveAddress = "0.0.0.0:8080"

func main() {
	// Get the gin web server runner
	apiRunner := gin.Default()

	// Load html pages
	apiRunner.LoadHTMLGlob("api/**/*.html")

	// Add routes
	apiRunner.GET("/basics/:input", basics.Basics)
	apiRunner.GET("/api/people", workshop_api.GetPeopleRoute)

	// Run the application on port 8080
	err := apiRunner.Run(serveAddress)

	// If an error occurs, fatally log the error
	fmt.Printf("Fatal error: %v", err.Error())
}
