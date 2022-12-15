package d14regolithreservoir

import (
	"advent/solutions/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const air = ' '
const rock = '#'
const sand = 'o'

const coordPattern = `([0-9]{0,3}\,[0-9]{0,3})`

type Cave struct {
	volume map[string]rune
	abyss  int
	floor  int
}

func (c *Cave) isEmpty(x, y int) bool {
	return c.volume[c.key(x, y)] == air
}

func (c *Cave) addLine(x1, y1, x2, y2 int) {
	c.setAbyss(y1, y2)
	if x1 == x2 {
		direction := 1
		if y2 < y1 {
			direction = -1
		}
		for i := y1; i != y2; i += direction {
			c.setVolume(x1, i, rock)
		}
		c.setVolume(x1, y2, rock)
		return
	}
	direction := 1
	if x2 < x1 {
		direction = -1
	}
	for i := x1; i != x2; i += direction {
		c.setVolume(i, y1, rock)
	}
	c.setVolume(x2, y1, rock)
}

func (c *Cave) setAbyss(y1, y2 int) {
	if y1 < y2 && y2 > c.abyss {
		c.abyss = y2
	}
	if y2 < y1 && y1 > c.abyss {
		c.abyss = y1
	}
	c.floor = c.abyss + 1
}

func (c *Cave) setVolume(x, y int, material rune) {
	c.volume[c.key(x, y)] = material
}

func (c *Cave) key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func (c *Cave) fillWithSand(sourceX, sourceY int, hasFloor bool) int {
	pastAbyss := false
	count := -1
	sourceKey := c.key(sourceX, sourceY)
	for !pastAbyss {
		pastAbyss = c.dropSand(sourceX, sourceY, hasFloor)
		count += 1
		if c.volume[sourceKey] == sand {
			break
		}
	}
	if hasFloor {
		count += 1
	}
	return count
}

func (c *Cave) dropSand(x, y int, hasFloor bool) bool {
	if y > c.abyss {
		if !hasFloor {
			return true
		}
		if y == c.floor {
			candidates := getCandidates(x, y)
			for _, coords := range candidates {
				c.setVolume(coords[0], coords[1], rock)
			}
			c.setVolume(x, y, sand)
			return false
		}
	}
	candidates := getCandidates(x, y)
	for _, coords := range candidates {
		_, ok := c.volume[c.key(coords[0], coords[1])]
		if !ok {
			return c.dropSand(coords[0], coords[1], hasFloor)
		}
	}
	c.setVolume(x, y, sand)
	return false
}

func (c *Cave) print() {
	vals := make([][]rune, 200)
	for i := range vals {
		vals[i] = make([]rune, 200)
		for j := range vals[i] {
			vals[i][j] = ' '
		}
	}
	for k, v := range c.volume {
		coords := parseCoords(k)
		vals[coords[0]-400][coords[1]] = v
	}
	for _, r := range vals {
		fmt.Println(string(r))
	}
}

func getCandidates(x, y int) [][]int {
	return [][]int{
		{x, y + 1},
		{x - 1, y + 1},
		{x + 1, y + 1},
	}
}

func FillCave(lines []string, hasFloor bool) int {
	c := &Cave{volume: map[string]rune{}}
	pattern := regexp.MustCompile(coordPattern)
	for _, l := range lines {
		matches := pattern.FindAllString(l, -1)
		prior := parseCoords(matches[0])
		for _, m := range matches[1:] {
			next := parseCoords(m)
			c.addLine(prior[0], prior[1], next[0], next[1])
			prior = next
		}
	}
	return c.fillWithSand(500, 0, hasFloor)
}

func parseCoords(s string) []int {
	components := strings.Split(s, ",")
	x, err := strconv.Atoi(components[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(components[1])
	if err != nil {
		log.Fatal(err)
	}
	return []int{x, y}
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(FillCave(lines, false)), strconv.Itoa(FillCave(lines, true))
}
