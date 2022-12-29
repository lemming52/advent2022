package d19notenoughminerals

import (
	"advent/solutions/utils"
	"fmt"
	"regexp"
	"strconv"
)

const blueprintPattern = `Blueprint (\d{1,2}): Each ore robot costs (\d{1,2}) ore. Each clay robot costs (\d{1,2}) ore. Each obsidian robot costs (\d{1,2}) ore and (\d{1,2}) clay. Each geode robot costs (\d{1,2}) ore and (\d{1,2}) obsidian.`
const maxTime = 24

type robot int

const (
	ore robot = iota
	clay
	obsidian
	geode
)

type BuildOrder struct {
	boundMin  int
	order     [][]int
	resources []int
	robots    []int
	counts    [][]int
	time      int
}

func nextBuildOrderStep(b *BuildOrder) *BuildOrder {
	new := &BuildOrder{
		boundMin:  0,
		order:     make([][]int, 24),
		resources: make([]int, 4),
		counts:    make([][]int, len(b.counts)),
		robots:    make([]int, 4),
		time:      b.time + 1,
	}
	copy(new.robots, b.robots)
	copy(new.resources, b.resources)
	copy(new.order, b.order)
	copy(new.counts, b.counts)
	return new
}

func (b *BuildOrder) calculateBound(scores []int) {
	b.boundMin = 0
	for i := ore; i <= geode; i++ {
		b.boundMin += b.robots[i] * scores[i]
	}
}

func (b *BuildOrder) generateOptions(costs [][]int) []*BuildOrder {
	options := []*BuildOrder{}
	options = append(options, nextBuildOrderStep(b))
	for i, c := range options[len(options)-1].robots {
		options[len(options)-1].resources[i] += c
	}
	for i := ore; i <= geode; i++ {
		buyCount := canAfford(b.resources, costs[i], int(i))
		if buyCount != 0 {
			new := nextBuildOrderStep(b)
			newBuildStep := make([]int, 4)
			newBuildStep[i] = buyCount
			new.order[new.time-1] = newBuildStep
			for j, c := range costs[i] {
				new.resources[j] -= c * buyCount
			}
			for j, c := range new.robots {
				new.resources[j] += c
			}
			new.robots[i] += buyCount
			options = append(options, new)
			if i == geode {
				return []*BuildOrder{new}
			}
		}
	}
	return options
}

func canAfford(resources, cost []int, rock int) int {
	for i, c := range cost {
		if c == 0 {
			continue
		}
		if resources[i] < c {
			return 0
		}
	}
	return 1
}

type Blueprint struct {
	name      int
	costs     [][]int
	robots    []int
	resources []int
	time      int
}

func newBlueprint(s string) *Blueprint {
	pattern := regexp.MustCompile(blueprintPattern)
	components := pattern.FindStringSubmatch(s)
	b := &Blueprint{
		name:      utils.Stoi(components[1]),
		costs:     make([][]int, 4),
		robots:    make([]int, 4),
		resources: make([]int, 4),
	}
	for i := ore; i <= geode; i++ {
		costs := make([]int, 3)
		switch i {
		case ore:
			costs[ore] = utils.Stoi(components[2])
			b.robots[ore] = 1
		case clay:
			costs[ore] = utils.Stoi(components[3])
			b.robots[clay] = 0
		case obsidian:
			costs[ore] = utils.Stoi(components[4])
			costs[clay] = utils.Stoi(components[5])
			b.robots[obsidian] = 0
		case geode:
			costs[ore] = utils.Stoi(components[6])
			costs[obsidian] = utils.Stoi(components[7])
			b.robots[clay] = 0
		}
		b.costs[i] = costs
		b.resources[i] = 0
	}
	return b
}

func (b *Blueprint) optimise() int {
	fmt.Println(b, "MARCO2")
	stack := newBuildStack()
	start := &BuildOrder{
		boundMin:  0,
		order:     make([][]int, 24),
		resources: make([]int, 4),
		robots:    make([]int, 4),
		time:      0,
	}
	copy(start.resources, b.resources)
	copy(start.robots, b.robots)
	stack.Push(start)
	best := start
	count := 0
	bestGeodes := make([]int, 25)
	for !stack.Empty() {
		count += 1
		current := stack.Pop().(*BuildOrder)
		if current.time == maxTime {
			if current.resources[geode] > bestGeodes[maxTime] {
				best = current
			}
		} else {
			iterations := current.generateOptions(b.costs)
			for _, i := range iterations {
				i.counts = append(i.counts, i.resources)
				if i.robots[geode] >= bestGeodes[i.time] {
					stack.Push(i)
				}
				if i.robots[geode] > bestGeodes[i.time] {
					bestGeodes[i.time] = i.robots[geode]
				}

			}
		}
	}
	fmt.Println(best.robots, best.resources, best.order, "FINAL")
	return best.resources[geode]
}

func BuildSCVs(lines []string) int {
	total := 0
	for _, l := range lines {
		b := newBlueprint(l)
		total += b.optimise()
	}
	return total
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(BuildSCVs(lines)), "B"
}
