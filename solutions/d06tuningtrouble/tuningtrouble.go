package d06tuningtrouble

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func scanForStart(stream string) int {
	prior := stream[0:3]
	count := 3
	i := 3
	for i < len(stream) {
		c := rune(stream[i])
		count += 1
		fmt.Println(count, i, string(c), prior)
		match := false
		for j, p := range prior {
			if p == c {
				match = true
				prior = stream[i+j-2 : i+j+1]
				i += j
				break
			}
		}
		if !match {
			return count
		}
	}
	return 0
}

func checkIfStart(s string, offset int) int {
	chars := map[rune]int{
		rune(s[0]): 0,
	}
	for i, c := range s[1:4] {
		if v, ok := chars[c]; ok {
			return checkIfStart(s[v+1:], offset+v+1)
		}
		chars[c] = i + 1
	}
	return offset + 4
}

func Run(path string) (string, string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return strconv.Itoa(checkIfStart(input, 0)), "B"
}
