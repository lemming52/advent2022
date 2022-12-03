package main

import (
	"advent/solutions/d01caloriecounting"
	"advent/solutions/d02rockpaperscissors"
	"advent/solutions/d03rucksackreorganisation"
	"flag"
	"fmt"
	"time"
)

func main() {
	var challenge string
	flag.StringVar(&challenge, "challenge", "sonarsweep", "name of challenge")
	all := flag.Bool("all", false, "display all results")
	flag.Parse()

	completed := []string{
		"caloriecounting",
		"rockpaperscissors",
		"rucksackreorganisation",
	}
	if *all {
		previous := time.Now()
		fmt.Println("Start Time: ", time.Now())
		for _, c := range completed {
			s := RunChallenge(c)
			current := time.Now()
			fmt.Println(s, " Duration/ms: ", float64(current.Sub(previous).Microseconds())/1000)
			previous = current
		}
	} else {
		fmt.Println(RunChallenge(challenge))
	}
}

func RunChallenge(challenge string) string {
	var res string
	switch challenge {
	case "caloriecounting", "1", "01":
        input := "inputs/d01caloriecounting.txt"
		A, B := d01caloriecounting.Run(input)
		res = fmt.Sprintf("caloriecounting Results A: %s B: %s", A, B)
	case "rockpaperscissors", "2", "02":
        input := "inputs/d02rockpaperscissors.txt"
		A, B := d02rockpaperscissors.Run(input)
		res = fmt.Sprintf("rockpaperscissors Results A: %s B: %s", A, B)
	case "rucksackreorganisation", "3", "03":
        input := "inputs/d03rucksackreorganisation.txt"
		A, B := d03rucksackreorganisation.Run(input)
		res = fmt.Sprintf("rucksackreorganisation Results A: %s B: %s", A, B)

    }
	return res
}
