package d12hillclimbingalgorithm

import (
	"advent/solutions/utils"
	"container/heap"
	"math"
	"strconv"
)

type Hill struct {
	height   int
	distance int
	index    int
	x, y     int
	visited  bool
}

func newHill(x, y, height int) *Hill {
	return &Hill{
		y:        y,
		x:        x,
		distance: math.MaxInt,
		visited:  false,
		height:   height,
	}
}

type Hillscape struct {
	hills [][]*Hill
	xMax  int
	yMax  int
	xGoal int
	yGoal int
}

func (h *Hillscape) dijkstra(x0, y0 int) {
	current := h.hills[x0][y0]
	current.distance = 0

	pq := PriorityQueue{current}
	heap.Init(&pq)

	for pq.Len() != 0 {
		current := heap.Pop(&pq).(*Hill)
		if current.x == h.xGoal && current.y == h.yGoal {
			break
		}
		current.visited = true
		neighbours := h.getNeighbours(current.x, current.y)
		for _, n := range neighbours {
			tentative := current.distance + 1
			if tentative < n.distance {
				n.distance = tentative
				heap.Push(&pq, n)
			}
		}
	}
}

func (h *Hillscape) getNeighbours(x, y int) []*Hill {
	current := h.hills[x][y]
	coords := utils.NESW2DNeighbours(x, y, h.xMax, h.yMax)
	correct := []*Hill{}
	for _, xy := range coords {
		if !h.hills[xy[0]][xy[1]].visited && h.hills[xy[0]][xy[1]].height <= current.height+1 {
			correct = append(correct, h.hills[xy[0]][xy[1]])
		}
	}
	return correct
}

func Navigate(input []string, variableStart bool) int {
	xStart, yStart := 0, 0
	xGoal, yGoal := 0, 0
	starts := [][]int{}
	for i, r := range input {
		for j, c := range r {
			switch c {
			case 'S':
				if variableStart {
					starts = append(starts, []int{i, j})
				} else {
					xStart = i
					yStart = j
				}
			case 'E':
				xGoal = i
				yGoal = j
			case 'a':
				if variableStart {
					starts = append(starts, []int{i, j})
				}
			}
		}
	}
	if variableStart {
		min := math.MaxInt
		for _, s := range starts {
			h := BuildHillscape(input, xGoal, yGoal)
			h.dijkstra(s[0], s[1])
			if h.hills[xGoal][yGoal].distance < min {
				min = h.hills[xGoal][yGoal].distance
			}
		}
		return min
	}
	hillscape := BuildHillscape(input, xGoal, yGoal)
	hillscape.dijkstra(xStart, yStart)
	return hillscape.hills[hillscape.xGoal][hillscape.yGoal].distance
}

func BuildHillscape(input []string, xGoal, yGoal int) *Hillscape {
	hillscape := &Hillscape{
		xMax:  len(input) - 1,
		yMax:  len(input[0]) - 1,
		xGoal: xGoal,
		yGoal: yGoal,
	}
	hills := make([][]*Hill, hillscape.xMax+1)
	for i := hillscape.xMax; i >= 0; i-- {
		hills[i] = make([]*Hill, (hillscape.yMax + 1))
		for j := hillscape.yMax; j >= 0; j-- {
			val := input[i][j]
			switch val {
			case 'S':
				val = 'a'
			case 'E':
				val = 'z'
			}
			n := newHill(i, j, int(val))
			hills[i][j] = n
		}
	}
	hillscape.hills = hills
	return hillscape
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(Navigate(lines, false)), strconv.Itoa(Navigate(lines, true))
}
