package d06tuningtrouble

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func findNonRepeatingSubstring(s string, offset, length int) int {
	chars := map[rune]int{
		rune(s[0]): 0,
	}
	for i, c := range s[1:length] {
		if v, ok := chars[c]; ok {
			return findNonRepeatingSubstring(s[v+1:], offset+v+1, length)
		}
		chars[c] = i + 1
	}
	return offset + length
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
	return strconv.Itoa(findNonRepeatingSubstring(input, 0, 4)), strconv.Itoa(findNonRepeatingSubstring(input, 0, 14))
}
