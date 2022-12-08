package d08treetoptreehouse

import (
	"advent/solutions/utils"
	"strconv"
)

const heightOffset = 48

type TreeMap struct {
	trees [][]int
	xMax  int
	yMax  int
}

func newTreeMap(lines []string) *TreeMap {
	trees := make([][]int, len(lines))
	for i, row := range lines {
		r := make([]int, len(row))
		for j, t := range row {
			r[j] = int(t - heightOffset)
		}
		trees[i] = r
	}
	return &TreeMap{
		trees: trees,
		xMax:  len(lines) - 1,
		yMax:  len(lines[0]) - 1,
	}
}

func (t *TreeMap) CountVisibleTrees() int {
	// account for edges
	count := 2*len(t.trees) + 2*len(t.trees[0]) - 4
	for i := 1; i < len(t.trees)-1; i++ {
		for j := 1; j < len(t.trees[0])-1; j++ {
			if t.treeIsVisible(i, j) {
				count += 1
			}
		}
	}
	return count
}

func (t *TreeMap) treeIsVisible(x, y int) bool {
	return t.visibleX(x, y, 0, x-1) ||
		t.visibleX(x, y, x+1, t.xMax) ||
		t.visibleY(x, y, 0, y-1) ||
		t.visibleY(x, y, y+1, t.yMax)
}

func (t *TreeMap) visibleX(x, y, xmin, xmax int) bool {
	for i := xmin; i <= xmax; i++ {
		if t.trees[i][y] >= t.trees[x][y] {
			return false
		}
	}
	return true
}

func (t *TreeMap) visibleY(x, y, ymin, ymax int) bool {
	for i := ymin; i <= ymax; i++ {
		if t.trees[x][i] >= t.trees[x][y] {
			return false
		}
	}
	return true
}

func DetermineSightlines(trees []string) int {
	treeMap := newTreeMap(trees)
	return treeMap.CountVisibleTrees()
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(DetermineSightlines(lines)), "B"
}
