package d18boilingboulders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanDroplet(t *testing.T) {
	input := []string{
		"2,2,2",
		"1,2,2",
		"3,2,2",
		"2,1,2",
		"2,3,2",
		"2,2,1",
		"2,2,3",
		"2,2,4",
		"2,2,6",
		"1,2,5",
		"3,2,5",
		"2,1,5",
		"2,3,5",
	}
	expected := 64
	res := ScanDroplet(input)
	assert.Equal(t, expected, res)
}
