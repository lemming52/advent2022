package d01caloriecounting

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func findMaxCalorie(values []int, target int) int {
	maxVals := make([]int, target)
	cumulative := 0
	for _, v := range values {
		if v == -1 {
			maxVals = updateMaxValues(maxVals, cumulative)
			cumulative = 0
			continue
		}
		cumulative += v
	}
	if cumulative != 0 {
		maxVals = updateMaxValues(maxVals, cumulative)
	}
	total := 0
	for _, v := range maxVals {
		total += v
	}
	return total
}

func updateMaxValues(maxVals []int, newVal int) []int {
	i := sort.SearchInts(maxVals, newVal)
	if i == len(maxVals) {
		maxVals = append(maxVals, newVal)
	} else {
		maxVals = append(maxVals[:i+1], maxVals[i:]...)
		maxVals[i] = newVal
	}
	return maxVals[1:]
}

func Run(path string) (string, string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	values := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			values = append(values, -1)
			continue
		}
		n, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, n)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return strconv.Itoa(findMaxCalorie(values, 1)), strconv.Itoa(findMaxCalorie(values, 3))
}
