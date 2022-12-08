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
	expected := 22
	res := DetermineSightlines(input)
	assert.Equal(t, expected, res)
}
