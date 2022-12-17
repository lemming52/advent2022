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

/*
func (c *Cave) exploreOptionWithElephant(name string, openValves map[string]bool, time, flowrate, total, busyMe, busyElephant int, me, elephant string) int {
	foundOption := false
	myOptions := []string{}
	elephantOptions := []string{}
	for _, n := range c.usefulValves {
		if ok := openValves[n]; ok {
			continue
		}
		newMyTime := time + busyMe + c.valves[me].distances[n].distance + 1
		if newMyTime <= c.maxTime {
			myOptions = append(myOptions, n)
		}
		newElephantTime := time + busyElephant + c.valves[elephant].distances[n].distance + 1
		if newElephantTime <= c.maxTime {
			elephantOptions = append(elephantOptions, n)
		}
	}
	options := combine(myOptions, elephantOptions)
	maxValue := total
	for _, option := range options {
		if option[0] == "pass" {


		} else if option[1] == "pass" {

		}
		foundOption = true
		newValves := make(map[string]bool, len(openValves)+2)
		for k, v := range openValves {
			newValves[k] = v
		}
		if busyMe == 0 && busyElephant == 0 {
			newValves[option[0]] = true
			newValves[option[1]] = true
			newMyTime := time + c.valves[me].distances[option[0]].distance + 1
			newElephantTime := time + c.valves[elephant].distances[option[0]].distance + 1

		}
		newTime :=
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
*/

func combinations(c []string) [][]string {
	res := [][]string{}
	for i, a := range c[:len(c)-1] {
		for _, b := range c[i+1:] {
			res = append(res, [][]string{{a, b}, {b, a}}...)
		}
	}
	return res
}

func combine(a, b []string) [][]string {
	res := [][]string{}
	if len(a) == 0 {
		for _, c := range b {
			res = append(res, []string{"pass", c})
		}
		return res
	}
	if len(b) == 0 {
		for _, c := range a {
			res = append(res, []string{c, "pass"})
		}
		return res
	}
	for _, c := range a {
		for _, d := range b {
			res = append(res, []string{c, d})
		}
	}
	return res
}

func (c *Cave) exploreElephantCrude() int {
	possibilities := []*Possibility{
		{
			me:           "AA",
			elephant:     "AA",
			busyMe:       0,
			busyElephant: 0,
			openValves:   map[string]int{},
		},
	}
	time := 4
	for time < c.maxTime {
		newPossibilities := []*Possibility{}
		maxScore := 0
		for _, p := range possibilities {
			if p.busyMe == 0 && p.busyElephant == 0 {
				foundOption := false
				for _, n := range combinations(c.usefulValves) {
					if _, ok := p.openValves[n[0]]; ok {
						continue
					}
					foundOption = true
					newMyTime := time + c.valves[p.me].distances[n[0]].distance
					if time == 4 {
						newMyTime += 1
					}
					meValid, elephantValid := false, false
					if newMyTime <= c.maxTime {
						meValid = true
					}
					newElephantTime := time + c.valves[p.elephant].distances[n[1]].distance
					if time == 4 {
						newElephantTime += 1
					}
					if newElephantTime <= c.maxTime {
						elephantValid = true
					}
					newValves := make(map[string]int, len(p.openValves)+2)
					for k, v := range p.openValves {
						newValves[k] = v
					}
					newPossibility := &Possibility{
						busyMe:       c.maxTime,
						busyElephant: c.maxTime,
						openValves:   newValves,
						me:           p.me,
						elephant:     p.elephant,
					}
					if meValid {
						newPossibility.busyMe = newMyTime - time
						newPossibility.me = n[0]
						newPossibility.openValves[n[0]] = newMyTime
					}
					if elephantValid {
						newPossibility.busyElephant = newElephantTime - time
						newPossibility.elephant = n[1]
						newPossibility.openValves[n[1]] = newElephantTime
					}
					score := score(c.valves, newPossibility.openValves, c.maxTime)
					if score > maxScore {
						maxScore = score
					}
					if score > maxScore*4/5 {
						newPossibilities = append(newPossibilities, newPossibility)
					}
				}
				if !foundOption {
					newValves := make(map[string]int, len(p.openValves)+2)
					for k, v := range p.openValves {
						newValves[k] = v
					}
					newPossibility := &Possibility{
						busyMe:       c.maxTime,
						busyElephant: c.maxTime,
						openValves:   newValves,
						me:           p.me,
						elephant:     p.elephant,
					}
					score := score(c.valves, newPossibility.openValves, c.maxTime)
					if score > maxScore {
						maxScore = score
					}
					if score > maxScore*4/5 {
						newPossibilities = append(newPossibilities, newPossibility)
					}
				}
			} else if p.busyMe == 0 {
				foundOption := false
				//fmt.Println(p.openValves)
				for _, n := range c.usefulValves {
					if _, ok := p.openValves[n]; ok {
						continue
					}
					//fmt.Print(n)
					newTime := time + c.valves[p.me].distances[n].distance
					newValves := make(map[string]int, len(p.openValves)+2)
					for k, v := range p.openValves {
						newValves[k] = v
					}
					var newPossibility *Possibility
					if newTime <= c.maxTime {
						newValves[n] = newTime
						newPossibility = &Possibility{
							busyMe:       newTime - time,
							busyElephant: p.busyElephant - 1,
							openValves:   newValves,
							me:           n,
							elephant:     p.elephant,
						}
					} else {
						newPossibility = &Possibility{
							busyMe:       c.maxTime,
							busyElephant: p.busyElephant - 1,
							openValves:   newValves,
							me:           p.me,
							elephant:     p.elephant,
						}
					}
					score := score(c.valves, newPossibility.openValves, c.maxTime)
					if score > maxScore {
						maxScore = score
					}
					if score > maxScore*4/5 {
						newPossibilities = append(newPossibilities, newPossibility)
					}
				}
				if !foundOption {
					newValves := make(map[string]int, len(p.openValves)+2)
					for k, v := range p.openValves {
						newValves[k] = v
					}
					newPossibility := &Possibility{
						busyMe:       c.maxTime,
						busyElephant: p.busyElephant - 1,
						openValves:   newValves,
						me:           p.me,
						elephant:     p.elephant,
					}
					score := score(c.valves, newPossibility.openValves, c.maxTime)
					if score > maxScore {
						maxScore = score
					}
					if score > maxScore*4/5 {
						newPossibilities = append(newPossibilities, newPossibility)
					}
				}
			} else if p.busyElephant == 0 {
				foundOption := false
				//if p.elephant == "DD" {
				///	fmt.Println("DD", p.busyElephant, p.busyMe, p.openValves, c.usefulValves)
				//}
				for _, n := range c.usefulValves {
					if _, ok := p.openValves[n]; ok {
						continue
					}
					foundOption = true
					newTime := time + c.valves[p.elephant].distances[n].distance
					newValves := make(map[string]int, len(p.openValves)+2)
					for k, v := range p.openValves {
						newValves[k] = v
					}
					//if p.elephant == "DD" && n == "HH" {
					//	fmt.Println(n, p.openValves, newTime)
					//}
					var newPossibility *Possibility
					if newTime <= c.maxTime {
						newValves[n] = newTime
						newPossibility = &Possibility{
							busyMe:       p.busyMe - 1,
							busyElephant: newTime - time,
							openValves:   newValves,
							me:           p.me,
							elephant:     n,
						}
					} else {
						newPossibility = &Possibility{
							busyMe:       p.busyMe - 1,
							busyElephant: c.maxTime,
							openValves:   newValves,
							me:           p.me,
							elephant:     p.elephant,
						}
					}
					score := score(c.valves, newPossibility.openValves, c.maxTime)
					if score > maxScore {
						maxScore = score
					}
					if score > maxScore*4/5 {
						newPossibilities = append(newPossibilities, newPossibility)
					}
				}
				if !foundOption {
					newValves := make(map[string]int, len(p.openValves)+2)
					for k, v := range p.openValves {
						newValves[k] = v
					}
					newPossibility := &Possibility{
						busyMe:       p.busyMe - 1,
						busyElephant: c.maxTime,
						openValves:   newValves,
						me:           p.me,
						elephant:     p.elephant,
					}
					score := score(c.valves, newPossibility.openValves, c.maxTime)
					if score > maxScore {
						maxScore = score
					}
					if score > maxScore*4/5 {
						newPossibilities = append(newPossibilities, newPossibility)
					}
				}
			} else {
				newValves := make(map[string]int, len(p.openValves)+2)
				for k, v := range p.openValves {
					newValves[k] = v
				}
				newPossibility := &Possibility{
					busyMe:       p.busyMe - 1,
					busyElephant: p.busyElephant - 1,
					openValves:   newValves,
					me:           p.me,
					elephant:     p.elephant,
				}
				newPossibilities = append(newPossibilities, newPossibility)
			}
		}
		time += 1
		possibilities = newPossibilities
	}
	maxFlow := 0
	for _, p := range possibilities {
		flow := 0
		for k, v := range p.openValves {
			flow += c.valves[k].flowrate * (c.maxTime - v)
		}
		if flow > maxFlow {
			maxFlow = flow
		}
	}
	return maxFlow
}

func score(valves map[string]*Valve, m map[string]int, maxTime int) int {
	flow := 0
	for k, v := range m {
		flow += valves[k].flowrate * (maxTime - v)
	}
	return flow
}

type Possibility struct {
	busyMe       int
	busyElephant int
	me           string
	elephant     string
	openValves   map[string]int
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

func SaveTheElephants(lines []string) (int, int) {
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
	return cave.exploreOption("AA", map[string]bool{}, 0, 0, 0, "AA"), cave.exploreElephantCrude()
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
	a, b := SaveTheElephants(lines)
	return strconv.Itoa(a), strconv.Itoa(b)
}
