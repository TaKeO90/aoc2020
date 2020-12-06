package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var (
	i          int
	singleTest []string
)

func getValues(inputs *[][]string, input string) {
	singleTest = append(singleTest, input)
	if i == 2 {
		(*inputs) = append((*inputs), singleTest)
		singleTest = []string{}
		i = 0
	} else {
		i++
	}
}

type passwordChecker struct {
	indecator string
	letter    string
	password  string
}

type passwordCheckers []passwordChecker

type solvingType int

const (
	Position solvingType = iota
	Iteration
)

func (ps passwordCheckers) solver(sType solvingType) int {
	var correctPasswords int
	for _, n := range ps {
		r1, r2 := strings.Split(n.indecator, "-")[0], strings.Split(n.indecator, "-")[1]
		ir1, _ := strconv.Atoi(r1)
		ir2, _ := strconv.Atoi(r2)
		switch sType {
		case Iteration:
			c := counter(n.letter, n.password)
			if c >= ir1 && c <= ir2 {
				correctPasswords++
			}
		case Position:
			pos := getPos(n.letter, n.password)
			ok1 := isIn(ir1, pos)
			ok2 := isIn(ir2, pos)
			if ok1 && !ok2 || !ok1 && ok2 {
				correctPasswords++
			}
		}
	}
	return correctPasswords
}

func counter(letter, password string) int {
	var counter int
	for _, i := range password {
		if string(i) == letter {
			counter++
		}
	}
	return counter
}

func isIn(pos int, positions []int) bool {
	for _, p := range positions {
		if p == pos {
			return true
		}
	}
	return false
}

func getPos(letter, password string) (positions []int) {
	for i, n := range password {
		if string(n) == letter {
			positions = append(positions, i+1)
		}
	}
	return
}

func newPasswordChecker(inputs [][]string) (pCheckers passwordCheckers) {
	pChecker := new(passwordChecker)
	for _, f := range inputs {
		for i, n := range f {
			switch i {
			case 0:
				pChecker.indecator = n
			case 1:
				pChecker.letter = strings.Split(n, ":")[0]
			case 2:
				pChecker.password = n
			}
		}
		pCheckers = append(pCheckers, *pChecker)
	}
	return
}

func parseFlags(p1, p2 *bool) {
	flag.BoolVar(p1, "partOne", false, "Choose solution for part 1")
	flag.BoolVar(p2, "partTwo", false, "Choose solution for part 2")
	flag.Parse()
}

func main() {
	var (
		input  string
		inputs [][]string
	)
	var (
		partOne bool
		partTwo bool
	)

	parseFlags(&partOne, &partTwo)

	for {
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			break
		}
		getValues(&inputs, input)
	}

	if partOne {
		fmt.Println(newPasswordChecker(inputs).solver(Iteration))
	} else if partTwo {
		fmt.Println(newPasswordChecker(inputs).solver(Position))
	} else {
		flag.PrintDefaults()
	}
}
