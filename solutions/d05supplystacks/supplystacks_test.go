package d05supplystacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	stacks := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
	}
	instructions := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}
	tests := []struct {
		name         string
		stacks       []string
		instructions []string
		part         bool
		expected     string
	}{
		{
			name:         "a",
			stacks:       stacks,
			instructions: instructions,
			part:         false,
			expected:     "CMZ",
		}, {
			name:         "b",
			stacks:       stacks,
			instructions: instructions,
			part:         true,
			expected:     "MCD",
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := execute(stacks, instructions, 3, tt.part)
			assert.Equal(t, tt.expected, res)
		})
	}

}
