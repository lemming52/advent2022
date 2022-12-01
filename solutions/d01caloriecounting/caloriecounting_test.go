package d01caloriecounting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxCalorie(t *testing.T) {
	input := []int{
		1000,
		2000,
		3000,
		-1,
		4000,
		-1,
		5000,
		6000,
		-1,
		7000,
		8000,
		9000,
		-1,
		10000,
	}
	expectedA := 24000
	expectedB := 45000
	resA := findMaxCalorie(input, 1)
	resB := findMaxCalorie(input, 3)
	assert.Equal(t, expectedA, resA)
	assert.Equal(t, expectedB, resB)
}
