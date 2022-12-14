package main

import (
	"advent/solutions/d01caloriecounting"
	"advent/solutions/d02rockpaperscissors"
	"advent/solutions/d03rucksackreorganisation"
	"advent/solutions/d04campcleanup"
	"advent/solutions/d05supplystacks"
	"advent/solutions/d06tuningtrouble"
	"advent/solutions/d07nospaceleftondevice"
	"advent/solutions/d08treetoptreehouse"
	"advent/solutions/d09ropebridge"
	"advent/solutions/d10cathoderaytube"
	"advent/solutions/d11monkeyinthemiddle"
	"advent/solutions/d12hillclimbingalgorithm"
	"advent/solutions/d13distresssignal"
	"flag"
	"fmt"
	"time"
)

func main() {
	var challenge string
	flag.StringVar(&challenge, "challenge", "campcleanup", "name or number of challenge")
	all := flag.Bool("all", false, "display all results")
	flag.Parse()

	completed := []string{
		"caloriecounting",
		"rockpaperscissors",
		"rucksackreorganisation",
		"campcleanup",
		"supplystacks",
		"tuningtrouble",
		"nospaceleftondevice",
		"treetoptreehouse",
		"ropebridge",
		"cathoderaytube",
		"monkeyinthemiddle",
		"hillclimbingalgorithm",
		"distresssignal",
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
	case "caloriecounting", "1":
        input := "inputs/d01caloriecounting.txt"
		A, B := d01caloriecounting.Run(input)
		res = fmt.Sprintf("caloriecounting Results A: %s B: %s", A, B)
	case "rockpaperscissors", "2":
        input := "inputs/d02rockpaperscissors.txt"
		A, B := d02rockpaperscissors.Run(input)
		res = fmt.Sprintf("rockpaperscissors Results A: %s B: %s", A, B)
	case "rucksackreorganisation", "3":
        input := "inputs/d03rucksackreorganisation.txt"
		A, B := d03rucksackreorganisation.Run(input)
		res = fmt.Sprintf("rucksackreorganisation Results A: %s B: %s", A, B)
	case "campcleanup", "4":
        input := "inputs/d04campcleanup.txt"
		A, B := d04campcleanup.Run(input)
		res = fmt.Sprintf("campcleanup Results A: %s B: %s", A, B)
	case "supplystacks", "5":
        input := "inputs/d05supplystacks.txt"
		A, B := d05supplystacks.Run(input)
		res = fmt.Sprintf("supplystacks Results A: %s B: %s", A, B)
	case "tuningtrouble", "6":
        input := "inputs/d06tuningtrouble.txt"
		A, B := d06tuningtrouble.Run(input)
		res = fmt.Sprintf("tuningtrouble Results A: %s B: %s", A, B)
	case "nospaceleftondevice", "7":
        input := "inputs/d07nospaceleftondevice.txt"
		A, B := d07nospaceleftondevice.Run(input)
		res = fmt.Sprintf("nospaceleftondevice Results A: %s B: %s", A, B)
	case "treetoptreehouse", "8":
        input := "inputs/d08treetoptreehouse.txt"
		A, B := d08treetoptreehouse.Run(input)
		res = fmt.Sprintf("treetoptreehouse Results A: %s B: %s", A, B)
	case "ropebridge", "9":
        input := "inputs/d09ropebridge.txt"
		A, B := d09ropebridge.Run(input)
		res = fmt.Sprintf("ropebridge Results A: %s B: %s", A, B)
	case "cathoderaytube", "10":
        input := "inputs/d10cathoderaytube.txt"
		A, B := d10cathoderaytube.Run(input)
		res = fmt.Sprintf("cathoderaytube Results A: %s B: %s", A, B)
	case "monkeyinthemiddle", "11":
        input := "inputs/d11monkeyinthemiddle.txt"
		A, B := d11monkeyinthemiddle.Run(input)
		res = fmt.Sprintf("monkeyinthemiddle Results A: %s B: %s", A, B)
	case "hillclimbingalgorithm", "12":
        input := "inputs/d12hillclimbingalgorithm.txt"
		A, B := d12hillclimbingalgorithm.Run(input)
		res = fmt.Sprintf("hillclimbingalgorithm Results A: %s B: %s", A, B)
	case "distresssignal", "13":
        input := "inputs/d13distresssignal.txt"
		A, B := d13distresssignal.Run(input)
		res = fmt.Sprintf("distresssignal Results A: %s B: %s", A, B)

    }
	return res
}
