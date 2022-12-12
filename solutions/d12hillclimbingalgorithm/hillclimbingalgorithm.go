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

func Navigate(input []string) int {
	hillscape := &Hillscape{
		xMax: len(input) - 1,
		yMax: len(input[0]) - 1,
	}
	hills := make([][]*Hill, hillscape.xMax+1)
	xStart, yStart := 0, 0
	for i := hillscape.xMax; i >= 0; i-- {
		hills[i] = make([]*Hill, (hillscape.yMax + 1))
		for j := hillscape.yMax; j >= 0; j-- {
			val := input[i][j]
			switch val {
			case 'S':
				val = 'a'
				xStart = i
				yStart = j
			case 'E':
				val = 'z'
				hillscape.xGoal = i
				hillscape.yGoal = j
			}
			n := newHill(i, j, int(val))
			hills[i][j] = n
		}
	}
	hillscape.hills = hills
	hillscape.dijkstra(xStart, yStart)
	return hillscape.hills[hillscape.xGoal][hillscape.yGoal].distance
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(Navigate(lines)), "B"
}
