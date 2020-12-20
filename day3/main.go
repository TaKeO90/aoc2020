package main

import (
	"fmt"
)

const tree = "#"

// PART 1
func traverseMap(Map [][]string) int {
	var (
		counter int
		x       int
	)
	for y, n := range Map {
		x = (x + 3) % len(n[0])
		if y != len(Map)-1 {
			if string(Map[y+1][0][x]) == tree {
				counter++
			}
		}
	}
	return counter
}

//PART 2
type Pattern int

const (
	RoDo Pattern = iota
	RtDo
	RfDo
	RsDo
	RoDt
)

func checkMultipleMapPaterns(pattern Pattern, Map [][]string) int {
	var (
		x       int
		counter int
		y2      int
	)

	for _, n := range Map {
		if y2 != len(Map)-1 {
			switch pattern {
			case RoDo:
				x = (x + 1) % len(n[0])
				y2++
			case RtDo:
				x = (x + 3) % len(n[0])
				y2++
			case RfDo:
				x = (x + 5) % len(n[0])
				y2++
			case RsDo:
				x = (x + 7) % len(n[0])
				y2++
			case RoDt:
				x = (x + 1) % len(n[0])
				y2 += 2
			}
			if string(Map[y2][0][x]) == tree {
				counter++
			}
		}
	}
	return counter
}

func main() {
	var (
		line  string
		lines [][]string
	)

	for {
		tmp := []string{}
		_, err := fmt.Scanf("%s", &line)
		if err != nil {
			break
		}
		tmp = append(tmp, line)
		lines = append(lines, tmp)
	}
	// IF PART ONE
	traverseMap(lines)
	// IF PART TWO
	var r int
	for i, n := range []Pattern{
		RoDo,
		RtDo,
		RfDo,
		RsDo,
		RoDt,
	} {
		result := checkMultipleMapPaterns(n, lines)
		if i == 0 {
			r = result
		} else {
			r *= result
		}
	}
	fmt.Println(r)
}
