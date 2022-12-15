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
	tests := []struct {
		name     string
		input    []string
		hasFloor bool
		expected int
	}{
		{
			name:     "a",
			input:    input,
			hasFloor: false,
			expected: 24,
		}, {
			name:     "b",
			input:    input,
			hasFloor: true,
			expected: 93,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := FillCave(tt.input, tt.hasFloor)
			assert.Equal(t, tt.expected, res)
		})
	}
}
