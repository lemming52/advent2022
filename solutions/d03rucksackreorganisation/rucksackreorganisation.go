package d03rucksackreorganisation

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const runeCutoff = 91
const lowerRuneOffset = 96
const capsRuneOffset = 38

func inspectRucksacks(rucksacks []string) int {
	score := 0
	for _, r := range rucksacks {
		item := findShared(r)
		score += itemPriority(item)
	}
	return score
}

func findShared(s string) rune {
	first := s[:len(s)/2]
	second := s[len(s)/2:]
	runes := map[rune]bool{}
	for _, r := range first {
		runes[r] = true
	}
	for _, r := range second {
		if _, ok := runes[r]; ok {
			return r
		}
	}
	return '0'
}

func itemPriority(r rune) int {
	if r > runeCutoff {
		return int(r - lowerRuneOffset)
	}
	return int(r - capsRuneOffset)
}

func Run(path string) (string, string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rucksacks := []string{}
	for scanner.Scan() {
		rucksacks = append(rucksacks, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return strconv.Itoa(inspectRucksacks(rucksacks)), "B"
}
