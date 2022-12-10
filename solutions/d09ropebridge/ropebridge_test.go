package d09ropebridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveRope(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		length   int
		expected int
	}{
		{
			name: "a",
			input: []string{
				"R 4",
				"U 4",
				"L 3",
				"D 1",
				"R 4",
				"D 1",
				"L 5",
				"R 2",
			},
			length:   2,
			expected: 13,
		}, {
			name: "b",
			input: []string{
				"R 5",
				"U 8",
				"L 8",
				"D 3",
				"R 17",
				"D 10",
				"L 25",
				"U 20",
			},
			length:   10,
			expected: 36,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := MoveRope(tt.input, tt.length)
			assert.Equal(t, tt.expected, res)
		})
	}
}
