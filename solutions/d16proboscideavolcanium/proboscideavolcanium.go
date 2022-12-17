package d16proboscideavolcanium

import (
	"advent/solutions/utils"
	"container/heap"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const valvePattern = `Valve ([A-Z]{2}) has flow rate=(\d{1,2}); (tunnels|tunnel) (lead|leads) to (valves|valve) ([A-Z\, ]+)`

type Cave struct {
	valves       map[string]*Valve
	usefulValves []string
	flowrate     int
	totalFlow    int
	time         int
	maxTime      int
}

func (c *Cave) escape(start string) int {
	current := c.valves[start]
	for c.time < c.maxTime {
		next := ""
		max := 0
		for k, v := range current.distances {
			if c.time+v.distance > c.maxTime {
				continue
			}
			if v.valve.open {
				continue
			}
			flowIncrease := v.valve.flowrate * (c.maxTime - c.time - v.distance - 1)
			if flowIncrease > max {
				next = k
				max = flowIncrease
			}
		}
		if next == "" {
			break
		}
		c.flow(current.distances[next].distance + 1)
		c.time += current.distances[next].distance + 1
		c.flowrate += current.distances[next].valve.flowrate
		current.distances[next].valve.open = true
		current = c.valves[next]
	}
	c.flow(c.maxTime - c.time)
	return c.totalFlow
}

func (c *Cave) exploreOption(name string, openValves map[string]bool, time, flowrate, total int, path string) int {
	foundOption := false
	maxValue := total
	for _, n := range c.usefulValves {
		if ok := openValves[n]; ok {
			continue
		}
		newTime := time + c.valves[name].distances[n].distance + 1
		if newTime > c.maxTime {
			continue
		}
		foundOption = true
		newValves := make(map[string]bool, len(openValves)+1)
		for k, v := range openValves {
			newValves[k] = v
		}
		newValves[n] = true
		result := c.exploreOption(
			n,
			newValves,
			newTime,
			flowrate+c.valves[n].flowrate,
			total+(flowrate*(c.valves[name].distances[n].distance+1)),
			fmt.Sprintf("%s,%s", path, n),
		)
		if result > maxValue {
			maxValue = result
		}
	}
	if !foundOption {
		maxValue += (c.maxTime - time) * flowrate
	}
	return maxValue
}

func (c *Cave) flow(time int) {
	c.totalFlow += c.flowrate * time
}

type Valve struct {
	name       string
	neighbours map[string]*Valve
	distances  map[string]*DistanceRecord
	flowrate   int
	open       bool
	closest    []string
}

type DistanceRecord struct {
	valve    *Valve
	distance int
	index    int
	visited  bool
	path     []string
}

func SaveTheElephants(lines []string) int {
	valves := BuildValves(lines)
	cave := &Cave{
		valves:       valves,
		flowrate:     0,
		totalFlow:    0,
		time:         0,
		maxTime:      30,
		usefulValves: []string{},
	}
	for k, v := range cave.valves {
		if v.flowrate > 0 {
			cave.usefulValves = append(cave.usefulValves, k)
		}
	}
	return cave.exploreOption("AA", map[string]bool{}, 0, 0, 0, "AA")
}

func BuildValves(lines []string) map[string]*Valve {
	pattern := regexp.MustCompile(valvePattern)
	valves := map[string]*Valve{}
	for _, l := range lines {
		components := pattern.FindStringSubmatch(l)
		v := &Valve{
			name:       components[1],
			flowrate:   utils.Stoi(components[2]),
			open:       false,
			neighbours: map[string]*Valve{},
		}
		neighbours := strings.Split(components[6], ",")
		for _, n := range neighbours {
			v.neighbours[strings.Trim(n, " ")] = nil
		}
		valves[v.name] = v
	}
	for name, v := range valves {
		for k := range v.neighbours {
			v.neighbours[k] = valves[k]
		}
		v.distances = make(map[string]*DistanceRecord, len(valves))
		v.distances[name] = &DistanceRecord{
			valve:    v,
			distance: 0,
			visited:  false,
			index:    0,
		}
		v.closest = make([]string, len(valves)-1)
		count := 0
		for k, val := range valves {
			if k == name {
				continue
			}
			v.distances[k] = &DistanceRecord{
				valve:    val,
				distance: math.MaxInt,
				index:    0,
				visited:  false,
			}
			v.closest[count] = k
			count += 1
		}
	}
	for _, v := range valves {
		ComputeDistances(v)
		sort.Slice(v.closest, func(i, j int) bool {
			return v.distances[v.closest[i]].distance < v.distances[v.closest[j]].distance
		})
	}
	return valves
}

func ComputeDistances(valve *Valve) {
	current := valve.distances[valve.name]
	pq := PriorityQueue{current}
	heap.Init(&pq)

	visitedCount := 0
	targetVisited := len(valve.distances)

	for pq.Len() != 0 {
		if visitedCount == targetVisited {
			break
		}
		current := heap.Pop(&pq).(*DistanceRecord)
		current.visited = true
		visitedCount += 1
		neighbours := current.valve.neighbours
		for _, n := range neighbours {
			tentative := current.distance + 1
			if tentative < valve.distances[n.name].distance {
				valve.distances[n.name].distance = tentative
				valve.distances[n.name].path = append(current.path, n.name)
				heap.Push(&pq, valve.distances[n.name])
			}
		}
	}
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(SaveTheElephants(lines)), "B"
}
