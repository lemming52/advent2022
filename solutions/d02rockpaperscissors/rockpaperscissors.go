package d02rockpaperscissors

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const runeOffset = 23
const runeMin = 88

func executeStrategyOne(games []string) int {
	scores := getScores()
	victors := getVictors()
	total := 0
	for _, g := range games {
		elems := strings.Split(g, " ")
		oppo, player := elems[0], elems[1]
		total += play(oppo, player, victors) + scores[player]
	}
	return total
}

func play(oppo, player string, victors map[string]string) int {
	if rune(oppo[0]) == rune(player[0])-runeOffset {
		return 3
	}
	if victors[player] == oppo {
		return 6
	}
	return 0
}

func getScores() map[string]int {
	return map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
}

func getVictors() map[string]string {
	return map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
	}
}

func executeStrategyTwo(games []string) int {
	scores := getScores()
	victors := getVictors()
	total := 0
	for _, g := range games {
		elems := strings.Split(g, " ")
		oppo, player := elems[0], elems[1]
		player = selectChoice(oppo, player)
		total += play(oppo, player, victors) + scores[player]
	}
	return total
}

func selectChoice(oppo, strategy string) string {
	new := '0'
	switch strategy {
	case "X":
		// loss
		new = rune(oppo[0]) + runeOffset - 1
	case "Y":
		// draw
		new = rune(oppo[0]) + runeOffset
	case "Z":
		// win
		new = rune(oppo[0]) + runeOffset + 1
	}
	if new < runeMin {
		new += 3
	} else if new > runeMin+2 {
		new -= 3
	}
	return string(new)
}

func Run(path string) (string, string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return strconv.Itoa(executeStrategyOne(lines)), strconv.Itoa(executeStrategyTwo(lines))
}
