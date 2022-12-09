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
	head           []int
	tail           []int
	positionsCount int
	relativeVector []int
	positionsMap   map[string]bool
}

func newMap() *KnotMap {
	positionsMap := map[string]bool{
		"0,0": true,
	}
	return &KnotMap{
		head:           []int{0, 0},
		tail:           []int{0, 0},
		positionsCount: 1,
		relativeVector: []int{0, 0},
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
		k.executeMove(moveVector)
	}
}

func (k *KnotMap) executeMove(move []int) {
	k.head[0] += move[0]
	k.head[1] += move[1]
	k.relativeVector[0] += move[0]
	k.relativeVector[1] += move[1]
	if k.moveTail() {
		k.examineTail()
	}
}

func (k *KnotMap) moveTail() bool {
	if magSquared(k.relativeVector[0], k.relativeVector[1]) <= ropeMagnitudeMax {
		return false
	}
	tailMove := []int{1, 1}
	for i, v := range k.relativeVector {
		if v < 0 {
			tailMove[i] = -1
		} else if v == 0 {
			tailMove[i] = 0
		}
	}
	k.tail[0] += tailMove[0]
	k.tail[1] += tailMove[1]
	k.relativeVector[0] = k.head[0] - k.tail[0]
	k.relativeVector[1] = k.head[1] - k.tail[1]
	return true
}

func magSquared(x, y int) int {
	return x*x + y*y
}

func (k *KnotMap) examineTail() {
	position := fmt.Sprintf("%d,%d", k.tail[0], k.tail[1])
	if ok := k.positionsMap[position]; !ok {
		k.positionsMap[position] = true
		k.positionsCount += 1
	}
}

func MoveRope(instructions []string) int {
	knots := newMap()
	for _, i := range instructions {
		knots.handleInstruction(i)
	}
	return knots.positionsCount
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(MoveRope(lines)), "B"
}
