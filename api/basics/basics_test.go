package basics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetResult_ReturnsEmptyString(t *testing.T) {
	// Act
	result, err := getResult("name")

	// Assert
	if assert.Nil(t, err) {
		assert.Equal(t, result, "")
	}
}
