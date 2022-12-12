# Advent of Code 2022 Solutions

My set of solutions for [Advent of Code 2022](https://adventofcode.com/2022), written in Golang because _I want practice_.

> In search of jungle star fruit.

## Running

My inputs from the challenges are all stored in the `inputs` directory, and at the time of writing these files are effectively hardcoded into the running.

To run a particular day (i.e. _campcleanup_, ...) use either the name or the day
```sh
go run main.go -challenge campcleanup
go run main.go -challenge 4
```

To run all days
```
go run main.go -all
```

### Challenge Days

Day | Challenge |Day | Challenge
----|-----------|----|----------
1 | `caloriecounting` | 14 | ` `
2 | `rockpaperscissors` | 15 | ` `
3 | `rucksackreorganisation` | 16 | ` `
4 | `campcleanup` | 17 | ` `
5 | `supplystacks` | 18 | ` `
6 | `tuningtrouble` | 19 | ` `
7 | `nospaceleftondevice` | 20 | ` `
8 | `treetoptreehouse` | 21 | ` `
9 | `ropebridge` | 22 | ` `
10 | `cathoderaytube` | 23 | ` `
11 | `monkeyinthemiddle` | 24 | ` `
12 | ` ` | 25 | ` `
13 | ` `

### Adding new template

To template out a new day, from the root directory run
```sh
python3 scripts/template.py <day_number> <challenge_name>
```
