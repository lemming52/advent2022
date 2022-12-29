package d19notenoughminerals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildSCVs(t *testing.T) {
	input := []string{
		"Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.",
		//"Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.",
	}
	expected := 33
	res := BuildSCVs(input)
	assert.Equal(t, expected, res)
}
