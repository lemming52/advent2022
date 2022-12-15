package d14regolithreservoir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFillCave(t *testing.T) {
	input := []string{
		"498,4 -> 498,6 -> 496,6",
		"503,4 -> 502,4 -> 502,9 -> 494,9",
	}
	expected := 24
	res := FillCave(input)
	assert.Equal(t, expected, res)
}
