package main

import (
	"flag"
	"fmt"
	"strings"
)

type Position struct {
	Row    int
	Column int
}

type Result int

var (
	results Results
)

type Results []Result

func (rs Results) Find(res Result) (int, bool) {
	for i, n := range rs {
		if res == n {
			return i, true
		}
	}
	return -1, false
}

func (r Result) Store() {
	results = append(results, r)
}

func getMax() (max Result) {
	for _, r := range results {
		if r > max {
			max = r
		}
	}
	return
}

func (p *Position) result() (r Result) {
	r = Result(p.Row)*8 + Result(p.Column)
	return
}

func getRow(min, max *int, index int, indicator string, pos *Position) {
	if index == 0 {
		switch indicator {
		case "F":
			*min = 0
			*max = 63
		case "B":
			*min = 64
			*max = 127
		}
	} else if index >= 1 && index <= 5 {
		switch indicator {
		case "F":
			*max = *min + (*max-*min)/2
		case "B":
			*min = (*min + (*max-*min)/2) + 1
		}
	} else if index == 6 {
		switch indicator {
		case "F":
			pos.Row = *min
		case "B":
			pos.Row = *max
		}
	}
}

func getColumn(index int, indicator string, columnMin, columnMax *int, pos *Position) {
	if index == 7 || index == 8 {
		switch indicator {
		case "R":
			*columnMin = (*columnMin + (*columnMax-*columnMin)/2) + 1
		case "L":
			*columnMax = *columnMin + (*columnMax-*columnMin)/2
		}
	} else {
		switch indicator {
		case "R":
			pos.Column = *columnMax
		case "L":
			pos.Column = *columnMin
		}
	}
}

func solve(input string, pos *Position) {
	pattern := strings.Split(input, "")
	var (
		max, min             int
		columnMin, columnMax int = 0, 7
	)

	for index, indicator := range pattern {
		getRow(&min, &max, index, indicator, pos)
		getColumn(index, indicator, &columnMin, &columnMax, pos)
	}
}

func solvePartI(input string) {
	var (
		pos Position
	)
	solve(input, &pos)
	if _, found := results.Find(pos.result()); !found {
		pos.result().Store()
	}
}

func solvePartII(input string) {
	var (
		pos Position
	)
	solve(input, &pos)
	if _, found := results.Find(pos.result()); !found {
		pos.result().Store()
	}
}

var (
	partI  bool
	partII bool
)

func main() {
	flag.BoolVar(&partI, "partI", false, "use it to solve part I")
	flag.BoolVar(&partII, "partII", false, "use it to solve part II")
	flag.Parse()
	var input string
	for {
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			break
		}
		if partI {
			solvePartI(input)
		} else if partII {
			solvePartII(input)
		} else {
			flag.PrintDefaults()
			break
		}

	}
	if len(results) != 0 && partI {
		fmt.Println(getMax())
		fmt.Println(results)
	} else if len(results) != 0 && partII {
		for r := 0; r < 128; r++ {
			for c := 0; c < 8; c++ {
				res := Result(r*8 + c)
				if _, found := results.Find(res); !found {
					_, foundF := results.Find(res - 1)
					_, foundS := results.Find(res + 1)
					if foundF && foundS {
						fmt.Println(res)
					}
				} else {
					continue
				}
			}
		}
	}
}
