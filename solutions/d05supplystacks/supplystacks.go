package d05supplystacks

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const movePattern = `move (\d+) from (\d+) to (\d+)`

// stackCount is the hardcoded number of columns used to read the data
const stackCount = 9

const spaceRune = rune(' ')

type Stacks struct {
	stacks      map[int][]rune
	stackCount  int
	movePattern *regexp.Regexp
}

func newStacks(stacks []string, count int) *Stacks {
	columns := map[int][]rune{}
	for i := 0; i < count; i++ {
		columns[i] = []rune{}
	}
	for i := len(stacks) - 1; i >= 0; i-- {
		row := stacks[i]
		for j := 0; j < count; j++ {
			val := rune(row[j*4+1])
			if val != spaceRune {
				columns[j] = append(columns[j], val)
			}
		}
	}
	pattern := regexp.MustCompile(movePattern)
	return &Stacks{
		stacks:      columns,
		stackCount:  count,
		movePattern: pattern,
	}
}

func (s *Stacks) execute(instruction string, moveMultiple bool) {
	components := s.movePattern.FindStringSubmatch(instruction)
	var err error
	values := make([]int, 3)
	for i := 1; i <= 3; i++ {
		values[i-1], err = strconv.Atoi(components[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	if moveMultiple {
		s.moveMultiple(values[0]-1, values[1]-1, values[2]-1)
	} else {
		for i := 0; i < values[0]; i++ {
			s.move(values[1]-1, values[2]-1)
		}
	}
}

func (s *Stacks) move(from, to int) {
	cutoff := len(s.stacks[from]) - 1
	s.stacks[to] = append(s.stacks[to], s.stacks[from][cutoff])
	s.stacks[from] = s.stacks[from][:cutoff]
}

func (s *Stacks) final() string {
	output := []rune{}
	for i := 0; i < s.stackCount; i++ {
		output = append(output, s.stacks[i][len(s.stacks[i])-1])
	}
	return string(output)
}

func (s *Stacks) print() {
	for i := 0; i < s.stackCount; i++ {
		fmt.Println(i, ": ", s.stacks[i])
	}
	fmt.Println()
}

func (s *Stacks) moveMultiple(count, from, to int) {
	cutoff := len(s.stacks[from]) - 1 - count
	s.stacks[to] = append(s.stacks[to], s.stacks[from][cutoff:]...)
	s.stacks[from] = s.stacks[from][:cutoff]
}

func execute(stacks, instructions []string, count int, partB bool) string {
	s := newStacks(stacks, count)
	for _, i := range instructions {
		s.execute(i, partB)
	}
	return s.final()
}

func Run(path string) (string, string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	stacks := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		stacks = append(stacks, text)
	}
	// drop index row
	stacks = stacks[:len(stacks)-1]

	instructions := []string{}
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return execute(stacks, instructions, stackCount, false), execute(stacks, instructions, stackCount, true)
}
