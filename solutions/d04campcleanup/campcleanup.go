package d04campcleanup

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

// 2-4,6-8
const assignmentPattern = `(\d+)\-(\d+),(\d+)\-(\d+)`

type Assignment struct {
	a0 int
	a1 int
	b0 int
	b1 int
}

func newAssignment(s string, pattern *regexp.Regexp) *Assignment {
	components := pattern.FindStringSubmatch(s)
	a := &Assignment{}
	for i, c := range components[1:] {
		val, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal(err)
		}
		switch i {
		case 0:
			a.a0 = val
		case 1:
			a.a1 = val
		case 2:
			a.b0 = val
		case 3:
			a.b1 = val
		}
	}
	return a
}

func (a *Assignment) fullyContained() bool {
	if a.a0 < a.b0 {
		return a.a1 >= a.b1
	}
	if a.b0 < a.a0 {
		return a.b1 >= a.a1
	}
	return true
}

func (a *Assignment) overlap() bool {
	if a.a0 < a.b0 {
		return a.b0 <= a.a1
	}
	if a.b0 < a.a0 {
		return a.a0 <= a.b1
	}
	return true
}

func spotOverlaps(assignments []string) (int, int) {
	pattern, err := regexp.Compile(assignmentPattern)
	if err != nil {
		log.Fatal(err)
	}
	contained, overlap := 0, 0
	for _, a := range assignments {
		assignment := newAssignment(a, pattern)
		if assignment.fullyContained() {
			contained += 1
			overlap += 1
		} else if assignment.overlap() {
			overlap += 1
		}
	}
	return contained, overlap
}

func Run(path string) (string, string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	assignments := []string{}
	for scanner.Scan() {
		assignments = append(assignments, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	a, b := spotOverlaps(assignments)
	return strconv.Itoa(a), strconv.Itoa(b)
}
