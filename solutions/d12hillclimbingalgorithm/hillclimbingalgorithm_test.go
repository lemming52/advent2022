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
	tests := []struct {
		name          string
		input         []string
		expected      int
		variableStart bool
	}{
		{
			name:          "a",
			input:         input,
			expected:      31,
			variableStart: false,
		}, {
			name:          "a",
			input:         input,
			expected:      29,
			variableStart: true,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := Navigate(tt.input, tt.variableStart)
			assert.Equal(t, tt.expected, res)
		})
	}

}
