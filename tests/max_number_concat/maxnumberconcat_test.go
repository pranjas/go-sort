package main

import (
	"pks_sort/algo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxNumberConcat(t *testing.T) {
	input := []string{"6", "66", "666", "6666"}
	maxNumber := algo.MaxNumber(input)
	t.Logf("Input = %v, maxNumber=%s", input, maxNumber)
	assert := assert.New(t)
	assert.Equal("6666666666", maxNumber, "Failed to get maximum number")
}

func TestMaxNumber2(t *testing.T) {
	input := []string{"123", "325", "6", "616", "623", "32", "11"}
	maxNumber := algo.MaxNumber(input)
	t.Logf("Input = %v, maxNumber=%s", input, maxNumber)
	assert := assert.New(t)
	assert.Equal("66236163253212311", maxNumber, "Failed to get maximum number")

}

func TestMaxNumber3(t *testing.T) {
	input := []string{"0", "8", "9", "6", "1", "2", "4", "3", "5", "7"}
	maxNumber := algo.MaxNumber(input)
	t.Logf("Input = %v, maxNumber=%s", input, maxNumber)
	assert := assert.New(t)
	assert.Equal("9876543210", maxNumber, "Failed to get maximum number")

}
