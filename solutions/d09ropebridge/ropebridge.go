package d09ropebridge

import (
	"advent/solutions/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const ropeMagnitudeMax = 2

type KnotMap struct {
	rope           [][]int
	positionsCount int
	tailIndex      int
	positionsMap   map[string]bool
}

func newMap(knotCount int) *KnotMap {
	positionsMap := map[string]bool{
		"0,0": true,
	}
	knots := make([][]int, knotCount)
	for i := 0; i < knotCount; i++ {
		knots[i] = make([]int, 2)
	}
	return &KnotMap{
		rope:           knots,
		positionsCount: 1,
		tailIndex:      knotCount - 1,
		positionsMap:   positionsMap,
	}
}

func (k *KnotMap) handleInstruction(s string) {
	moveVector := make([]int, 2)
	components := strings.Split(s, " ")
	distance, err := strconv.Atoi(components[1])
	if err != nil {
		log.Fatal(err)
	}
	switch components[0] {
	case "U":
		moveVector[0], moveVector[1] = 0, 1
	case "D":
		moveVector[0], moveVector[1] = 0, -1
	case "L":
		moveVector[0], moveVector[1] = -1, 0
	case "R":
		moveVector[0], moveVector[1] = 1, 0
	}
	for i := 0; i < distance; i++ {
		k.rope[0][0] += moveVector[0]
		k.rope[0][1] += moveVector[1]
		nextIndex := 1
		movedNext := k.moveKnot(nextIndex)
		for movedNext && nextIndex < k.tailIndex {
			nextIndex += 1
			movedNext = k.moveKnot(nextIndex)
		}
		k.examineTail()
	}
}

func (k *KnotMap) moveKnot(currentIndex int) bool {
	relativeVector := make([]int, 2)
	relativeVector[0] = k.rope[currentIndex-1][0] - k.rope[currentIndex][0]
	relativeVector[1] = k.rope[currentIndex-1][1] - k.rope[currentIndex][1]
	if magSquared(relativeVector[0], relativeVector[1]) <= ropeMagnitudeMax {
		return false
	}
	tailMove := []int{1, 1}
	for i, v := range relativeVector {
		if v < 0 {
			tailMove[i] = -1
		} else if v == 0 {
			tailMove[i] = 0
		}
	}
	k.rope[currentIndex][0] += tailMove[0]
	k.rope[currentIndex][1] += tailMove[1]
	return true
}

func magSquared(x, y int) int {
	return x*x + y*y
}

func (k *KnotMap) examineTail() {
	position := fmt.Sprintf("%d,%d", k.rope[k.tailIndex][0], k.rope[k.tailIndex][1])
	if ok := k.positionsMap[position]; !ok {
		k.positionsMap[position] = true
		k.positionsCount += 1
	}
}

func MoveRope(instructions []string, ropeLength int) int {
	knots := newMap(ropeLength)
	for _, i := range instructions {
		knots.handleInstruction(i)
	}
	return knots.positionsCount
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(MoveRope(lines, 2)), strconv.Itoa(MoveRope(lines, 10))
}
