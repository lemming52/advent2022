package main

import (
	"advent/solutions/d01sonarsweep"
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
		"sonarsweep",
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
	input := fmt.Sprintf("inputs/%s.txt", challenge)
	switch challenge {
	case "sonarsweep", "1", "01":
		A, B := d01sonarsweep.LoadSonar(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	}
	return res
}
