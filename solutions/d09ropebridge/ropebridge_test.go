package d09ropebridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveRope(t *testing.T) {
	input := []string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2",
	}
	expected := 13
	res := MoveRope(input)
	assert.Equal(t, expected, res)
}
