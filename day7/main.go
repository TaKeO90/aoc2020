package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strings"
)

func mustGetFileContent(filename string) (content []string) {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	return
}

type Bag struct {
	Color   string
	Content string
	Amount  int
}

type Bags []Bag

func parseContent(content []string) Bags {
	getBag := func(line string) string {
		bag := strings.TrimSpace(strings.Split(line, "bags")[0])
		return bag
	}

	getBagContent := func(line string) []string {
		var contents []string
		bagContents := strings.Split(line, "contain")[1]
		if strings.Contains(bagContents, ",") {
			for _, c := range strings.Split(bagContents, ",") {
				contents = append(contents, strings.TrimSpace(c))
			}
		} else {
			contents = append(contents, bagContents)
		}
		return contents
	}

	var bags Bags
	var noBags string = "no other bags."

	singleBag := new(Bag)
	for _, line := range content {
		if bagCnt := getBagContent(line); len(bagCnt) >= 1 {
			for _, bCnt := range bagCnt {
				if !strings.Contains(bCnt, noBags) {
					var amount int
					var containedBag, bColor, ignore string
					singleBag.Color = getBag(line)
					fmt.Sscanf(bCnt, "%d %s %s %s", &amount, &containedBag, &bColor, &ignore)
					singleBag.Content = containedBag + " " + bColor
					singleBag.Amount = amount
					bags = append(bags, *singleBag)
					singleBag = &Bag{}
				}
			}
		}
	}
	return bags
}

func (b Bags) Solve() {
	startColor := "shiny gold"
	fmt.Println(CountBagsV1(b, startColor))
	fmt.Println(CountBagsV2(b, startColor))
}

func CountBagsV1(bags Bags, color string) int {
	lst := list.New()
	visited := map[string]bool{}

	lst.PushBack(color)
	for lst.Len() > 0 {
		next := lst.Front()
		if _, ok := visited[next.Value.(string)]; !ok {
			for _, c := range getContainersByColor(bags, next.Value.(string)) {
				lst.PushBack(c)
			}
		}
		visited[next.Value.(string)] = true
		lst.Remove(next)
	}
	return len(visited) - 1
}

func getContainersByColor(bags Bags, color string) (result []string) {
	for _, bag := range bags {
		if bag.Content == color {
			result = append(result, bag.Color)
		}
	}
	return
}

func CountBagsV2(bags Bags, color string) (masterSum int) {
	for _, b := range bags {
		if b.Color == color {
			masterSum += b.Amount + b.Amount*CountBagsV2(bags, b.Content)
		}
	}
	return
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("no input file\n")
		os.Exit(1)
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("recovered from %v\n", r)
		}
	}()

	for _, f := range os.Args[1:] {
		content := mustGetFileContent(f)
		bags := parseContent(content)
		bags.Solve()
	}
}
