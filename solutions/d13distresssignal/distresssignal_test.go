package d13distresssignal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseElements(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "[1,2,3]",
			expected: []string{"1", "2", "3"},
		}, {
			input:    "[[1],[2,3,4]]",
			expected: []string{"[1]", "[2,3,4]"},
		}, {
			input:    "[1,[2,[3,[4,[5,6,7]]]],8,9]",
			expected: []string{"1", "[2,[3,[4,[5,6,7]]]]", "8", "9"},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.input, func(t *testing.T) {
			res := parseElements(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestPacketOrderFull(t *testing.T) {
	input := []string{
		"[1,1,3,1,1]",
		"[1,1,5,1,1]",
		"",
		"[[1],[2,3,4]]",
		"[[1],4]",
		"",
		"[9]",
		"[[8,7,6]]",
		"",
		"[[4,4],4,4]",
		"[[4,4],4,4,4]",
		"",
		"[7,7,7,7]",
		"[7,7,7]",
		"",
		"[]",
		"[3]",
		"",
		"[[[]]]",
		"[[]]",
		"",
		"[1,[2,[3,[4,[5,6,7]]]],8,9]",
		"[1,[2,[3,[4,[5,6,0]]]],8,9]",
	}
	expected := 13
	res := PacketOrder(input)
	assert.Equal(t, expected, res)
}

func TestPacketOrder(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
				"[[4,10,[[8],[6,4,8,8,10],1,[10,2],[7,8,9,0]],1],[6,4,3],[7,8,4],[[9,[10,4,8,7,2]]],[3,5]]",
				"[[],[[[0,0,5,10],9,10],[[6,7,5]],[[3,6,7,2,2],[9],[7,5],9],[[0,0,10,9,8],5,[3,2]],[]]]",
			},
			expected: 0,
		}, {
			input: []string{
				"[[[]],[]]",
				"[[8,[5,[6,3,2],[10,5,4],[5,1],[5,10,5,9,7]],4,[8,2,0,6],4],[[[0,9,5,7],9],6,[[0,8,10,8],9,[9,0,3,7,1],[6,8,2,5,10],[]],7,[1,[],4,[0],[2]]],[1],[[7,[],6,2,1]],[2,3]]",
			},
			expected: 1,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.input[0], func(t *testing.T) {
			res := PacketOrder(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}
