package main

import (
	"advent/solutions/d01caloriecounting"
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
		res = fmt.Sprintf("%s Results A: %s B: %s", challenge, A, B)

	}
	return res
}
