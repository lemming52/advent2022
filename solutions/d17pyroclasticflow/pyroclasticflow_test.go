package d17pyroclasticflow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRockFlow(t *testing.T) {
	input := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	expected := 3068
	res := RockFlow(input)
	assert.Equal(t, expected, res)
}
