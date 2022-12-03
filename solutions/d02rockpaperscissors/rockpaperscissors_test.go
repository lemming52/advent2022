package d02rockpaperscissors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteStrategyOne(t *testing.T) {
	input := []string{
		"A Y",
		"B X",
		"C Z",
	}
	expected := 15
	res := executeStrategyOne(input)
	assert.Equal(t, expected, res)
}

func TestExecuteStrategyTwo(t *testing.T) {
	input := []string{
		"A Y",
		"B X",
		"C Z",
	}
	expected := 12
	res := executeStrategyTwo(input)
	assert.Equal(t, expected, res)
}

func TestPlay(t *testing.T) {
	victors := getVictors()
	tests := []struct {
		oppo     string
		player   string
		expected int
	}{
		{
			oppo:     "A",
			player:   "Y",
			expected: 6,
		}, {
			oppo:     "B",
			player:   "X",
			expected: 0,
		}, {
			oppo:     "C",
			player:   "Z",
			expected: 3,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.oppo, func(t *testing.T) {
			res := play(tt.oppo, tt.player, victors)
			assert.Equal(t, tt.expected, res, "returned value should match expected	")
		})
	}
}
