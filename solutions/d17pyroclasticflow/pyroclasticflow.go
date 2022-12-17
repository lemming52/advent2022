package d17pyroclasticflow

import (
	"advent/solutions/utils"
	"fmt"
	"strconv"
)

// 7 empty spaces for empty row
const emptyRow = "       "

const spawnOffset = 4

type Cave struct {
	rows   [][]rune
	time   int
	height int
	width  int
}

func (c *Cave) spawnRock() ([][]int, int) {
	c.time += 1
	switch c.time % 5 {
	case 0:
		return [][]int{
			{0, 2},
			{0, 3},
			{0, 4},
			{0, 5},
		}, 1
	case 1:
		return [][]int{
			{0, 3},
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 3},
		}, 3
	case 2:
		return [][]int{
			{0, 2},
			{0, 3},
			{0, 4},
			{1, 4},
			{2, 4},
		}, 3
	case 3:
		return [][]int{
			{0, 2},
			{1, 2},
			{2, 2},
			{3, 2},
		}, 4
	case 4:
		return [][]int{
			{0, 2},
			{0, 3},
			{1, 2},
			{1, 3},
		}, 2
	default:
		return nil, 0
	}
}

func (c *Cave) handleRock(jets string, index int) int {
	rock, height := c.spawnRock()
	requiredHeight := c.height + spawnOffset + height
	if len(c.rows) < requiredHeight {
		c.rows = append(c.rows, emptyRows(requiredHeight-len(c.rows))...)
	}
	horizontalOffset := 0
	verticalOffset := c.height + spawnOffset
	for true {
		shift := 0
		switch jets[index] {
		case '<':
			shift = -1
		case '>':
			shift = 1
		}
		index += 1
		if index == len(jets) {
			index = 0
		}
		if !c.willCollide(rock, horizontalOffset+shift, verticalOffset) {
			horizontalOffset += shift
		}
		if c.willCollide(rock, horizontalOffset, verticalOffset-1) {
			break
		}
		verticalOffset -= 1
	}
	for _, r := range rock {
		c.rows[r[0]+verticalOffset][r[1]+horizontalOffset] = '#'
		if r[0]+verticalOffset > c.height {
			c.height = r[0] + verticalOffset
		}
	}
	return index
}

func (c *Cave) willCollide(rock [][]int, hOffset, vOffset int) bool {
	for _, v := range rock {
		x, y := v[0]+vOffset, v[1]+hOffset
		if x < 0 || y >= c.width || y < 0 {
			return true
		}
		if c.rows[x][y] == '#' {
			return true
		}
	}
	return false
}

func (c *Cave) print() {
	fmt.Println(c.time, c.height)
	for i := len(c.rows) - 1; i >= 0; i-- {
		fmt.Println(string(c.rows[i]))
	}
}

func emptyRows(h int) [][]rune {
	rows := make([][]rune, h)
	for i := range rows {
		rows[i] = []rune(emptyRow)
	}
	return rows
}

func RockFlow(jets string) int {
	c := Cave{
		rows:  [][]rune{[]rune("#######")},
		time:  -1,
		width: 7,
	}
	index := 0
	for c.time < 2021 {
		index = c.handleRock(jets, index)
	}
	return c.height
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(RockFlow(lines[0])), "B"
}
