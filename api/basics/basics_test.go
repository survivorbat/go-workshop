package basics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// The test method
func Test_GetResult_ReturnsExpectedResult(t *testing.T) {
	// Our test data, we provide both a name (the key) and
	// a struct that we create on the spot. It doesn't matter
	// what you call the key (hello, good, tacocat) as long as
	// they are unique.
	// This struct has 2 properties:
	//
	// input -> What we're going to give the getResult function
	// expectedOutput -> What we expect to get in return, so JSON
	//
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

	// Use range to loop through the key-values, name is the key
	// and testData is the struct
	for name, testData := range tests {
		// Run the test using the name
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

// The test method
func Test_GetResult_ReturnsErrorOnSingleCharacter(t *testing.T) {
	tests := map[string]struct{
		input string
		expectedResult string
	}{
		// First test data set, hello should return olleh
		"a": {input: "a", expectedResult: "{\"error\":\"input should be longer than 1 character\"}"},

		// Second test data set, good should return doog
		"b": {input: "b", expectedResult: "{\"error\":\"input should be longer than 1 character\"}"},
	}

	// Use range to loop through the key-values, name is the key
	// and testData is the struct
	for name, testData := range tests {
		// Run the test using the name
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
