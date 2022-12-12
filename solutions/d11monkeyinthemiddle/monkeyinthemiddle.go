package d11monkeyinthemiddle

import (
	"advent/solutions/utils"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const operationPattern = `  Operation: new = old ([+\-*\/]{1}) ([a-z0-9]+)`
const targetPattern = `    If (true|false): throw to monkey (\d)+`
const testPattern = `  Test: divisible by ([0-9]+)`

type OperationFunction func(int) int
type TestFunction func(int) bool

type Monkey struct {
	items         []int
	name          int
	operationFunc OperationFunction
	testFunc      TestFunction
	divisor       int
	trueTarget    int
	falseTarget   int
	inspectCount  int
}

func (m *Monkey) inspect(targets map[int]*Monkey, skipDivision bool, lcd int) {
	for _, i := range m.items {
		m.inspectCount += 1
		worry := m.operationFunc(i)
		if !skipDivision {
			worry = int(worry / 3)
		}
		target := 0
		if m.testFunc(worry) {
			target = m.trueTarget
		} else {
			target = m.falseTarget
		}
		t, ok := targets[target]
		if !ok {
			log.Fatal("no target", target)
		}
		worry = worry % lcd
		t.items = append(t.items, worry)
	}
	m.items = []int{}
}

func ParseMonkey(input []string, number int) *Monkey {
	monkey := &Monkey{
		name:         number,
		inspectCount: 0,
	}
	ParseItemLine(input[1], monkey)
	ParseOperationLine(input[2], monkey)
	ParseTestLines(input[3:6], monkey)
	return monkey
}

func ParseItemLine(s string, m *Monkey) {
	items := strings.Split(s, " ")[4:]
	for _, i := range items {
		val, err := strconv.Atoi(strings.Trim(i, ","))
		if err != nil {
			log.Fatal(err)
		}
		m.items = append(m.items, val)
	}
}

func ParseOperationLine(s string, m *Monkey) {
	pattern := regexp.MustCompile(operationPattern)
	components := pattern.FindStringSubmatch(s)
	useInput := components[2] == "old"
	val := 0
	var err error
	if !useInput {
		val, err = strconv.Atoi(components[2])
		if err != nil {
			log.Fatal(err)
		}
	}
	switch components[1] {
	case "+":
		m.operationFunc = Addition(val, useInput)
	case "*":
		m.operationFunc = Multiplication(val, useInput)
	case "-", "/":
		panic("not implemented")
	}
}

func Addition(factor int, useInput bool) OperationFunction {
	if useInput {
		return func(input int) int {
			return input + input
		}
	} else {
		return func(input int) int {
			return input + factor
		}
	}
}

func Multiplication(factor int, useInput bool) OperationFunction {
	if useInput {
		return func(input int) int {
			return input * input
		}
	} else {
		return func(input int) int {
			return input * factor
		}
	}
}

func ParseTestLines(lines []string, m *Monkey) {
	testLine := lines[0]
	trueTarget, falseTarget := lines[1], lines[2]
	testP := regexp.MustCompile(testPattern)
	targetP := regexp.MustCompile(targetPattern)
	testComponents := testP.FindStringSubmatch(testLine)
	divisibleVal, err := strconv.Atoi(testComponents[1])
	if err != nil {
		log.Fatal(err)
	}
	m.testFunc = Divisible(divisibleVal)
	m.divisor = divisibleVal
	trueComponents := targetP.FindStringSubmatch(trueTarget)
	trueVal, err := strconv.Atoi(trueComponents[2])
	m.trueTarget = trueVal
	falseComponents := targetP.FindStringSubmatch(falseTarget)
	falseVal, err := strconv.Atoi(falseComponents[2])
	m.falseTarget = falseVal
}

func Divisible(factor int) TestFunction {
	return func(input int) bool {
		return input%factor == 0
	}
}

func MonkeyBusiness(input []string, rounds int, skipDivision bool) int {
	monkeys := map[int]*Monkey{}
	for i := 0; i <= len(input)/7; i++ {
		monkeys[i] = ParseMonkey(input[i*7:], i)
	}
	Play(monkeys, rounds, skipDivision)
	values := make([]int, len(monkeys))
	for k, v := range monkeys {
		values[k] = v.inspectCount
	}
	sort.Ints(values)
	return values[len(values)-1] * values[len(values)-2]
}

func Play(monkeys map[int]*Monkey, rounds int, skipDivision bool) {
	lcd := 1
	for _, v := range monkeys {
		lcd *= v.divisor
	}
	for i := 0; i < rounds; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkeys[j].inspect(monkeys, skipDivision, lcd)
		}
	}
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(MonkeyBusiness(lines, 20, false)), strconv.Itoa(MonkeyBusiness(lines, 10000, true))
}
