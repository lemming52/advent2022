package d08treetoptreehouse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetermineSightlines(t *testing.T) {
	input := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}
	expectedA, expectedB := 21, 8
	resA, resB := DetermineSightlines(input)
	assert.Equal(t, expectedA, resA)
	assert.Equal(t, expectedB, resB)

}
