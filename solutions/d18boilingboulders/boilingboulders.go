package d18boilingboulders

import (
	"advent/solutions/utils"
	"fmt"
	"strconv"
	"strings"
)

type Cube struct {
	name   string
	coords []int
	faces  map[string]*Cube
}

func newCube(s string) *Cube {
	coords := make([]int, 3)
	for i, v := range strings.Split(s, ",") {
		coords[i] = utils.Stoi(v)
	}
	c := &Cube{
		name:   key(coords),
		coords: coords,
		faces:  make(map[string]*Cube, 6),
	}
	c.generateFaceNeighbours()
	return c
}

func (c *Cube) generateFaceNeighbours() {
	x, y, z := c.coords[0], c.coords[1], c.coords[2]
	neighbours := [][]int{
		{x - 1, y, z},
		{x + 1, y, z},
		{x, y - 1, z},
		{x, y + 1, z},
		{x, y, z - 1},
		{x, y, z + 1},
	}
	for _, n := range neighbours {
		c.faces[key(n)] = nil
	}
}

func (c *Cube) addNeighbours(library map[string]*Cube) {
	for k := range c.faces {
		if v, ok := library[k]; ok {
			c.faces[k] = v
			v.attachFace(c)
		}
	}
}

func (c *Cube) attachFace(neighbour *Cube) {
	c.faces[neighbour.name] = neighbour
}

func (c *Cube) surfaceArea() int {
	sa := 6
	for _, v := range c.faces {
		if v != nil {
			sa -= 1
		}
	}
	return sa
}

func key(c []int) string {
	return fmt.Sprintf("%d,%d,%d", c[0], c[1], c[2])
}

func ScanDroplet(lines []string) int {
	library := map[string]*Cube{}
	for _, l := range lines {
		c := newCube(l)
		library[c.name] = c
		c.addNeighbours(library)
	}
	total := 0
	for _, v := range library {
		total += v.surfaceArea()
	}
	return total
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(ScanDroplet(lines)), "B"
}
