package d01caloriecounting

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func FindMaxCalorie(values []int) int {
	maxVal, cumulative := 0, 0
	for _, v := range values {
		if v == -1 {
			if cumulative > maxVal {
				maxVal = cumulative
			}
			cumulative = 0
			continue
		}
		cumulative += v
	}
	return maxVal
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
	return strconv.Itoa(FindMaxCalorie(values)), "B"
}
