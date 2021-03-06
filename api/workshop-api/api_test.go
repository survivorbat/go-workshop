package workshop_api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// For the getPeople function

func TestGetPeople_ReturnsExpectedData(t *testing.T) {
	// Arrange
	expectedData := "not implemented"

	// Act
	result, err := getPeople()

	// Assert
	assert.Nil(t, result)
	assert.Equal(t, expectedData, err.Error())
}

// For the GetPeopleRoute function

func TestGetPeopleRoute_ReturnsExpectedData(t *testing.T) {
	// Arrange
	expectedJson := "{}"

	// Response will be written to this writer
	writer := httptest.NewRecorder()

	// Test context for Gin
	c, _ := gin.CreateTestContext(writer)

	// Add a request to the context, so we can extract values from it later (title)
	c.Request = httptest.NewRequest(http.MethodGet, "https://example.com", nil)

	// Act
	GetPeopleRoute(c)

	// Assert
	// TODO: Add status code and body check, remove print statement
	fmt.Print(expectedJson)
}
