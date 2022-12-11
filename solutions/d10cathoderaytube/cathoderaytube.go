package d10cathoderaytube

import (
	"advent/solutions/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type CathodeCPU struct {
	register     int
	cycleNumber  int
	outputSum    int
	outputRows   []string
	renderingRow []rune
}

func newCPU() *CathodeCPU {

	return &CathodeCPU{
		register:     1,
		cycleNumber:  0,
		outputSum:    0,
		outputRows:   []string{},
		renderingRow: newEmptyRow(),
	}
}

func (c *CathodeCPU) cycle(count int) {
	for i := 0; i < count; i++ {
		c.cycleNumber += 1
		rowPosition := (c.cycleNumber - 1) % 40
		if rowPosition == 0 && c.cycleNumber != 1 {
			c.outputRows = append(c.outputRows, string(c.renderingRow))
			c.renderingRow = newEmptyRow()
		}
		if rowPosition >= c.register-1 && rowPosition <= c.register+1 {
			c.renderingRow[rowPosition] = '#'
		}
		if (c.cycleNumber-20)%40 == 0 {
			c.outputSum += c.cycleNumber * c.register
		}
	}
}

func (c *CathodeCPU) instruct(s string) {
	switch s[0] {
	case 'n':
		c.cycle(1)
	case 'a':
		c.cycle(2)
		components := strings.Split(s, " ")
		val, err := strconv.Atoi(components[1])
		if err != nil {
			log.Fatal(err)
		}
		c.register += val
	}
}

func newEmptyRow() []rune {
	renderingRow := make([]rune, 40)
	for i := range renderingRow {
		renderingRow[i] = '.'
	}
	return renderingRow
}

func Execute(instructions []string) (int, []string) {
	cpu := newCPU()
	for _, i := range instructions {
		cpu.instruct(i)
	}
	cpu.outputRows = append(cpu.outputRows, string(cpu.renderingRow))
	return cpu.outputSum, cpu.outputRows
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	val, output := Execute(lines)
	for _, r := range output {
		fmt.Println(r)
	}
	return strconv.Itoa(val), "PAPKFKEJ"
}
