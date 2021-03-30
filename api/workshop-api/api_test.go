package workshop_api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
Exercise 2
*/

//// For the getPeople function
//
//func TestGetPeople_ReturnsExpectedData(t *testing.T) {
//	// Arrange
//	expectedData := People
//
//	// Act
//	result, err := getPeople()
//
//	// Assert
//	assert.Nil(t, err)
//	assert.Equal(t, expectedData, result)
//}
//
//// For the GetPeopleRoute function
//// We don't have multiple data sets because we're returning static data
//
//func TestGetPeopleRoute_ReturnsExpectedData(t *testing.T) {
//	// Arrange
//	expectedJson := "[{\"name\":\"Lindsay\",\"title\":\"Developer\"},{\"name\":\"Bob\",\"title\":\"Product Owner\"},{\"name\":\"Chris\",\"title\":\"Developer\"},{\"name\":\"Dahlia\",\"title\":\"Operations\"}]"
//
//	// Response will be written to this writer
//	writer := httptest.NewRecorder()
//
//	// Test context for Gin
//	c, _ := gin.CreateTestContext(writer)
//
//	// Act
//	GetPeopleRoute(c)
//
//	// Assert
//	assert.Equal(t, 200, writer.Code)
//	assert.Equal(t, expectedJson, writer.Body.String())
//}

/**
Exercise 3
*/

func TestGetPeople_ReturnsExpectedData(t *testing.T) {
	tests := map[string]struct{
		inputTitle string
		expectedPeople []string
	}{
		// Use your own data here ;-)
		"developer": {
			inputTitle: "Developer",
			expectedPeople: []string{"Lindsay", "Chris"},
		},
		"product owner": {
			inputTitle: "Product Owner",
			expectedPeople: []string{"Bob"},
		},
	}

	for name, testData := range tests {
		t.Run(name, func(t *testing.T) {
			// Act
			result, err := getPeople(testData.inputTitle)

			// Assert
			assert.Nil(t, err)

			// Check if they're the same length
			assert.Equal(t, len(testData.expectedPeople), len(result))

			for _, person := range result {
				assert.Contains(t, testData.expectedPeople, person.Name)
			}
		})
	}
}
