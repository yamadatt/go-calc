package utils_test

import (
	"testing"

	"go-calc/utils" // Replace with the correct import path

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	actual := utils.Sum(3, 5)
	expected := 8

	assert.Equal(t, expected, actual)
}

func TestSumNegativeNumbers(t *testing.T) {
	actual := utils.Sum(-3, -5)
	expected := -8

	assert.Equal(t, expected, actual)
}

func TestSumZero(t *testing.T) {
	actual := utils.Sum(0, 0)
	expected := 0

	assert.Equal(t, expected, actual)
}

func TestSumLargeNumbers(t *testing.T) {
	actual := utils.Sum(1000000, 2000000)
	expected := 3000000

	assert.Equal(t, expected, actual)
}
