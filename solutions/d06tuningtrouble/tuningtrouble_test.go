package d06tuningtrouble

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIfStart(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: 7,
		},
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 5,
		}, {
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 6,
		}, {
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 10,
		}, {
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 11,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.input, func(t *testing.T) {
			res := checkIfStart(tt.input, 0)
			assert.Equal(t, tt.expected, res)
		})
	}

}
