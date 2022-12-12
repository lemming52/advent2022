package d12hillclimbingalgorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNavigate(t *testing.T) {
	input := []string{
		"Sabqponm",
		"abcryxxl",
		"accszExk",
		"acctuvwj",
		"abdefghi",
	}
	expected := 31
	res := Navigate(input)
	assert.Equal(t, expected, res)
}
