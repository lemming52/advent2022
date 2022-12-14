package d13distresssignal

import (
	"advent/solutions/utils"
	"fmt"
	"log"
	"strconv"
)

func compare(l, r string) (bool, bool) {
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
	return lVal < rVal, lVal != rVal
}

func compareLists(l, r string, leftConverted, rightConverted bool) (bool, bool) {
	if len(l) == 2 {
		return true, len(r) != 2
	}
	if len(r) == 2 {
		return false, true
	}
	elemL, elemR := parseElements(l), parseElements(r)
	if len(elemL) == 0 {
		return true, len(elemR) != 0
	}
	for i, e := range elemL {
		if i >= len(elemR) {
			return false, true
		}
		correct, decided := compare(e, elemR[i])
		if decided {
			return correct, true
		}
	}
	return true, len(elemL) != len(elemR)
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
		val, _ := compare(lines[i*3], lines[i*3+1])
		if val {
			sum += i + 1
		}
	}
	return sum
}

func PacketSort(lines []string) int {
	packets := []string{}
	for i := 0; i <= len(lines)/3; i++ {
		packets = append(packets, lines[i*3:i*3+2]...)
	}
	packets = append(packets, []string{"[[2]]", "[[6]]"}...)
	for i := 0; i < len(packets)-1; i++ {
		for j := len(packets) - 1; j > i; j-- {
			if correct, _ := compare(packets[j-1], packets[j]); !correct {
				packets[j], packets[j-1] = packets[j-1], packets[j]
			}
		}
	}
	product := 1
	for i, v := range packets {
		if v == "[[2]]" || v == "[[6]]" {
			product *= (i + 1)
		}
	}
	return product
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	return strconv.Itoa(PacketOrder(lines)), strconv.Itoa(PacketSort(lines))
}
