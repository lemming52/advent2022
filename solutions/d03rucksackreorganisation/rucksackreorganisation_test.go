package d03rucksackreorganisation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInspectRucksacks(t *testing.T) {
	input := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
	expected := 157
	res := inspectRucksacks(input)
	assert.Equal(t, expected, res)
}

func TestItemPriority(t *testing.T) {
	tests := []struct {
		r        rune
		expected int
	}{
		{
			r:        'p',
			expected: 16,
		}, {
			r:        'L',
			expected: 38,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(string(tt.r), func(t *testing.T) {
			res := itemPriority(tt.r)
			assert.Equal(t, tt.expected, res, "returned value should match expected	")
		})
	}
}
