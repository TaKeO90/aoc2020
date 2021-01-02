package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	m       = map[string]int{}
	result  int
	count   int
	lastElm bool
)

// SolveFile Solve the challenge for each individual file.
func SolveFile(filePath string) {
	data := mustReadFile(filePath)
	splitedData := strings.Split(data, "\n")

	var (
		anyOneResult   int
		everyOneResult int
	)

	// First Part
	for _, n := range splitedData {
		getGroupAnswers(n, &anyOneResult, anyone)
	}

	// Second Part
	answers := [][]string{}
	gAnsw := strings.Split(data, "\n")
	tmp := []string{}
	for _, n := range gAnsw {
		if n != "" {
			tmp = append(tmp, n)
		} else {
			if len(tmp) != 0 {
				answers = append(answers, tmp)
				tmp = []string{}
			}
		}
	}
	everyOneResult = getGroupAnswer(answers)

	fmt.Println("anyOneResult", anyOneResult)
	fmt.Println("everyOneResult", everyOneResult)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("USAGE: ./day6 <input files>\n")
		os.Exit(1)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered from %v\n", r)
		}
	}()

	files := os.Args[1:]

	for _, filePath := range files {
		SolveFile(filePath)
		reset()
	}
}

func reset() {
	m = map[string]int{}
	result = 0
	count = 0
	lastElm = false
}

type answerType int

const (
	anyone answerType = iota
	everyone
)

func getGroupAnswers(answers string, fResult *int, aType answerType) {
	if answers != "" {
		count++
		groupAnswer := strings.Split(answers, "")
		for _, ans := range groupAnswer {
			m[ans] = 1
		}
	} else {
		for _, v := range m {
			result += v
		}
		*fResult += result
		reset()
	}
}

func mustReadFile(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func getGroupAnswer(groupAnswers [][]string) (numberOfAnwsers int) {
	isSingle := false
	for _, a := range groupAnswers {
		answCounter := map[string]int{}
		if len(a) > 1 {
			isSingle = false
			for _, n := range a {
				if len(n) > 1 {
					nn := strings.Split(n, "")
					for _, f := range nn {
						answCounter[f] += 1
					}
				} else {
					answCounter[n] += 1
				}
			}
		} else {
			isSingle = true
			for _, g := range a {
				if len(g) > 1 {
					for _, h := range g {
						answCounter[string(h)] += 1
					}
				} else {
					answCounter[g] += 1
				}
			}
		}
		if isSingle {
			numberOfAnwsers += getFinalResult(answCounter, single, len(a))
		} else {
			r := getFinalResult(answCounter, multiple, len(a))
			numberOfAnwsers += r
		}
	}
	return
}

type SingleOrMultiple int

const (
	single SingleOrMultiple = iota
	multiple
)

func getFinalResult(m map[string]int, sOm SingleOrMultiple, length int) int {
	switch sOm {
	case single:
		var r int
		for _, n := range m {
			r += n
		}
		return r
	case multiple:
		var r int
		var tmp []int
		for _, n := range m {
			tmp = append(tmp, n)
		}
		r = areAllEqual(tmp, length)
		return r
	default:
		return 0
	}
}

func areAllEqual(arr []int, length int) int {
	r := 0
	for _, n := range arr {
		if n == length {
			r++
		}
	}
	return r
}
