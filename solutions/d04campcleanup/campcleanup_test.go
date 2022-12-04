package d04campcleanup

import (
	"log"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpotOverlaps(t *testing.T) {
	input := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}
	expected := 2
	res := spotOverlaps(input)
	assert.Equal(t, expected, res)
}

func TestNewAssignment(t *testing.T) {
	tests := []struct {
		s        string
		expected *Assignment
	}{
		{
			s: "2-4,6-8",
			expected: &Assignment{
				a0: 2,
				a1: 4,
				b0: 6,
				b1: 8,
			},
		},
	}
	pattern, err := regexp.Compile(assignmentPattern)
	if err != nil {
		log.Fatal(err)
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.s, func(t *testing.T) {
			res := newAssignment(tt.s, pattern)
			assert.Equal(t, tt.expected, res, "returned value should match expected	")
		})
	}
}

func TestFullyContained(t *testing.T) {

}
