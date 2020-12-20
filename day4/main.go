package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	Fields = map[string]bool{
		"ecl": true,
		"pid": true,
		"eyr": true,
		"hcl": true,
		"byr": true,
		"iyr": true,
		"cid": true,
		"hgt": true,
	}
)

func passportChecker(target [][]string) {
	var (
		val     int
		counter int
	)
	for _, p := range target {
		pp := strings.Split(strings.Join(p, " "), " ")
		val = checkSinglePassport(pp)
		if val == 8 {
			counter++
		}
	}
	fmt.Println(counter)
}

func checkSinglePassport(passport []string) int {
	var (
		validation int
		isCid      bool
		keys       []string
	)
	for _, n := range passport {
		if n != "" {
			key := strings.Split(n, ":")[0]
			value := strings.Split(strings.TrimSpace(n), ":")[1]
			keys = append(keys, key)
			isValueok := valuesChecks(key, value)
			ok := Fields[key]
			if ok && isValueok {
				validation++
			}
		}
	}
	if validation < 8 {
		for _, k := range keys {
			if k == "cid" {
				isCid = true
				break
			}
		}
		if !isCid {
			validation++
		}
	}
	return validation
}

func valuesChecks(key, value string) bool {
	switch key {
	case "byr":
		var v int
		fmt.Sscanf(value, "%d", &v)
		if v >= 1920 && v <= 2002 {
			return true
		}
	case "iyr":
		var v int
		fmt.Sscanf(value, "%d", &v)
		if v >= 2010 && v <= 2020 {
			return true
		}
	case "eyr":
		var v int
		fmt.Sscanf(value, "%d", &v)
		if v >= 2020 && v <= 2030 {
			return true
		}
	case "hgt":
		var (
			n int
			v string
		)
		fmt.Sscanf(value, "%d%s", &n, &v)
		if v == "cm" {
			if n >= 150 && n <= 193 {
				return true
			}
		} else if v == "in" {
			if n >= 59 && n <= 76 {
				return true
			}
		}
	case "hcl":
		ok, err := regexp.Match(`[0-9]|[a-f]`, []byte(value))
		if err != nil {
			return false
		}
		if strings.HasPrefix(value, "#") && ok {
			return true
		}
	case "ecl":
		switch value {
		case "amb":
			return true
		case "blu":
			return true
		case "brn":
			return true
		case "gry":
			return true
		case "grn":
			return true
		case "hzl":
			return true
		case "oth":
			return true
		default:
			return false
		}
	case "pid":
		if len(value) == 9 {
			return true
		}
	case "cid":
		return true
	}
	return false
}

func main() {
	var (
		targets [][]string
	)
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	splittedData := strings.Split(string(data), "\n")
	tmp := []string{}
	for _, d := range splittedData {
		tmp = append(tmp, strings.TrimSpace(d))
		if d == "" {
			targets = append(targets, tmp)
			tmp = []string{}
		}
	}
	//fmt.Println(targets)
	passportChecker(targets)
}
