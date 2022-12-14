package d13distresssignal

import (
	"advent/solutions/utils"
	"fmt"
	"log"
	"strconv"
)

func compare(l, r string) bool {
	//fmt.Println(l, r, "compare")
	leftList, rightList := l[0] == '[', r[0] == '['
	if leftList {
		if rightList {
			return compareLists(l, r, false, false)
		}
		return compareLists(l, wrapList(r), false, true)
	}
	if rightList {
		return compareLists(wrapList(l), r, true, false)
	}
	lVal, err := strconv.Atoi(l)
	if err != nil {
		log.Fatal(err)
	}
	rVal, err := strconv.Atoi(r)
	if err != nil {
		log.Fatal(err)
	}
	return lVal <= rVal
}

func compareLists(l, r string, leftConverted, rightConverted bool) bool {
	//fmt.Println(l, r, "compareLists")
	if len(l) == 2 {
		return true
	}
	if len(r) == 2 {
		return false
	}
	elemL, elemR := parseElements(l), parseElements(r)
	if !rightConverted && !leftConverted && len(elemR) < len(elemL) {
		return false
	}
	if len(elemL) == 0 {
		return true
	}
	for i, e := range elemL {
		if i >= len(elemR) {
			return rightConverted
		}
		if !compare(e, elemR[i]) {
			return false
		}
	}
	return true
}

func parseElements(s string) []string {
	s = s[1 : len(s)-1]
	openCount := 0
	currentElementIndex := 0
	elements := []string{}
	for i, c := range s {
		switch c {
		case '[':
			openCount += 1
		case ']':
			openCount -= 1
		case ',':
			if openCount == 0 {
				elements = append(elements, s[currentElementIndex:i])
				currentElementIndex = i + 1
			}
		}
	}
	elements = append(elements, s[currentElementIndex:])
	return elements
}

func wrapList(s string) string {
	return fmt.Sprintf("[%s]", s)
}

func PacketOrder(lines []string) int {
	sum := 0
	for i := 0; i <= len(lines)/3; i++ {
		val := compare(lines[i*3], lines[i*3+1])
		if val {
			sum += i + 1
		}
		fmt.Println(i, lines[i*3], lines[i*3+1], val)
	}
	return sum
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(PacketOrder(lines)), "B"
}
