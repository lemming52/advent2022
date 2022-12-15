package d15beaconexclusionzone

import (
	"advent/solutions/utils"
	"log"
	"regexp"
	"sort"
	"strconv"
)

const beaconPattern = `Sensor at x=([\-]{0,1}\d+)\, y=([\-]{0,1}\d+): closest beacon is at x=([\-]{0,1}\d+)\, y=([\-]{0,1}\d+)`

type Beacon struct {
	x, y               int
	closestX, closestY int
	radius             int
}

func newBeacon(coords []int) *Beacon {
	b := &Beacon{
		x:        coords[0],
		y:        coords[1],
		closestX: coords[2],
		closestY: coords[3],
	}
	b.calculateRadius()
	return b
}

func (b *Beacon) calculateRadius() {
	deltaX := mag(b.x, b.closestX)
	deltaY := mag(b.y, b.closestY)
	b.radius = deltaX + deltaY
}

func (b *Beacon) rowExclusion(y int) []int {
	deltaY := mag(b.y, y)
	width := b.radius - deltaY
	if width < 0 {
		return nil
	}
	return []int{b.x - width, b.x + width}
}

func rowExclusion(beacons []*Beacon, y int) int {
	ranges := [][]int{}
	for _, b := range beacons {
		r := b.rowExclusion(y)
		if r != nil {
			ranges = append(ranges, r)
		}
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	prior := ranges[0]
	excluded := 0
	for _, r := range ranges[1:] {
		if r[0] <= prior[1] {
			if r[1] > prior[1] {
				prior[1] = r[1]
			}
		} else {
			excluded += mag(prior[1], prior[0])
			prior = r
		}
	}
	excluded += mag(prior[1], prior[0])
	return excluded
}

func frequency(beacons []*Beacon, min, max int) int {
	for y := min; y <= max; y++ {
		ranges := [][]int{}
		for _, b := range beacons {
			r := b.rowExclusion(y)
			if r != nil {
				ranges = append(ranges, r)
			}
		}
		sort.Slice(ranges, func(i, j int) bool {
			return ranges[i][0] < ranges[j][0]
		})
		prior := []int{ranges[0][0], ranges[0][1]}
		frequency := -1
		for _, r := range ranges[1:] {
			if r[0] <= prior[1] {
				if r[1] > prior[1] {
					prior[1] = r[1]
				}
				continue
			}
			if r[0] <= min {
				copy(prior, r)
				continue
			}
			if r[0] >= max {
				break
			}
			frequency = (r[0]-1)*4000000 + y
			break
		}
		if frequency != -1 {
			return frequency
		}
	}
	return 0
}

func ExclusionZone(lines []string, y int) int {
	pattern := regexp.MustCompile(beaconPattern)
	beacons := []*Beacon{}
	for _, l := range lines {
		components := pattern.FindStringSubmatch(l)
		coords := make([]int, 4)
		for i, c := range components[1:] {
			coords[i] = stoi(c)
		}
		beacons = append(beacons, newBeacon(coords))
	}
	return rowExclusion(beacons, y)
}

func Distress(lines []string, x, y int) int {
	pattern := regexp.MustCompile(beaconPattern)
	beacons := []*Beacon{}
	for _, l := range lines {
		components := pattern.FindStringSubmatch(l)
		coords := make([]int, 4)
		for i, c := range components[1:] {
			coords[i] = stoi(c)
		}
		beacons = append(beacons, newBeacon(coords))
	}
	return frequency(beacons, x, y)
}

func stoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func mag(a, b int) int {
	delta := a - b
	if delta < 0 {
		delta *= -1
	}
	return delta
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(ExclusionZone(lines, 2000000)), strconv.Itoa(Distress(lines, 0, 4000000))
}
