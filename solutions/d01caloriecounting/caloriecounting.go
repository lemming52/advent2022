package d01caloriecounting

import (
	"bufio"
	"encoding/json"
	"fmt"
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

type Entry struct {
	Val        int   `json:"val"`
	MaxIndices []int `json:"maxIndices"`
}

func findMaxCaloriePrint(values []int, target int) (int, []*Entry) {
	maxVals := make([]int, target)
	cumulative := 0
	entries := []*Entry{}
	indices := map[int]int{}
	counter := -1
	for _, v := range values {
		if v == -1 {
			counter += 1
			indices[cumulative] = counter
			maxVals = updateMaxValues(maxVals, cumulative)
			index := make([]int, target)
			for i, v := range maxVals {
				index[i] = indices[v]
			}
			entries = append(entries, &Entry{
				Val:        cumulative,
				MaxIndices: index,
			})
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
	return total, entries
}

func contains(sl []int, t int) bool {
	for _, value := range sl {
		if value == t {
			return true
		}
	}
	return false
}

func writeEntries(path string, entries []*Entry) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintln(f, "[")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries[:len(entries)-1] {
		out, err := json.Marshal(e)
		if err != nil {
			log.Fatal(err)
		}
		_, err = fmt.Fprintln(f, fmt.Sprintf("\t%s,", out))
		if err != nil {
			log.Fatal(err)
		}
	}
	out, err := json.Marshal(entries[len(entries)-1])
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintln(f, fmt.Sprintf("\t%s", out))
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintln(f, "]")
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func RunPrint(path, output string) (string, string) {
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
	a, _ := findMaxCaloriePrint(values, 1)
	b, entries := findMaxCaloriePrint(values, 3)
	writeEntries(output, entries)
	return strconv.Itoa(a), strconv.Itoa(b)
}
