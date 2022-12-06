package d06tuningtrouble

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindNonRepeatingSubstring(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		expected int
	}{
		{
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			length:   4,
			expected: 7,
		},
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			length:   4,
			expected: 5,
		}, {
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			length:   4,
			expected: 6,
		}, {
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			length:   4,
			expected: 10,
		}, {
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			length:   4,
			expected: 11,
		}, {
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			length:   14,
			expected: 19,
		},
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			length:   14,
			expected: 23,
		}, {
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			length:   14,
			expected: 23,
		}, {
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			length:   14,
			expected: 29,
		}, {
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			length:   14,
			expected: 26,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.input, func(t *testing.T) {
			res := findNonRepeatingSubstring(tt.input, 0, tt.length)
			assert.Equal(t, tt.expected, res)
		})
	}

}
