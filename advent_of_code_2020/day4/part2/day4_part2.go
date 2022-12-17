package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const VALID = "byr,iyr,eyr,hgt,hcl,ecl,pid"

func main() {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	r := 0
	for _, s := range strings.Split(string(f), "\n\n") {
		if checkValid(s) {
			if checkAttributes(buildMap((s))) {
				r++
			}
		}
	}
	fmt.Println(r)
}

func checkValid(s string) bool {
	for _, n := range strings.Split(VALID, ",") {
		if !strings.Contains(s, n) {
			return false
		}
	}
	return true
}

func buildMap(s string) map[string]string {
	m := make(map[string]string)
	for _, sl := range strings.Split(string(s), "\n") {
		for _, str := range strings.Split(sl, " ") {
			sl := strings.Split(str, ":")
			m[sl[0]] = sl[1]
		}
	}
	return m
}

/*
byr (Birth Year) - four digits; at least 1920 and at most 2002.
iyr (Issue Year) - four digits; at least 2010 and at most 2020.
eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
hgt (Height) - a number followed by either cm or in:
If cm, the number must be at least 150 and at most 193.
If in, the number must be at least 59 and at most 76.
hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
pid (Passport ID) - a nine-digit number, including leading zeroes.
cid (Country ID) - ignored, missing or not.
*/

func checkAttributes(m map[string]string) bool {
	for key := range m {
		switch key {
		case "byr":
			v, _ := strconv.Atoi(m[key])
			if v < 1920 || v > 2002 {
				return false
			}
		case "iyr":
			v, _ := strconv.Atoi(m[key])
			if v < 2010 || v > 2020 {
				return false
			}
		case "eyr":
			v, _ := strconv.Atoi(m[key])
			if v < 2020 || v > 2030 {
				return false
			}
		case "hgt":
			reg, err := regexp.Compile("[^0-9]+")
			if err != nil {
				log.Panic(err)
			}
			numstr, _ := strconv.Atoi(reg.ReplaceAllString(m[key], ""))
			if strings.Contains(m[key], "cm") {
				if numstr < 150 || numstr > 193 {
					return false
				}
			} else if strings.Contains(m[key], "in") {
				if numstr < 59 || numstr > 76 {
					return false
				}
			} else if !strings.Contains(m[key], "cm") && !strings.Contains(m[key], "in") {
				return false
			}
		case "hcl":
			b, err := regexp.MatchString("#[0-9a-f]{6}", m[key])
			if err != nil {
				log.Panic(err)
			}
			if !b {
				return false
			}
		case "ecl":
			b, err := regexp.Match(m[key], []byte("\\b(amb|blu|brn|gry|grn|hzl|oth)\\b"))
			if err != nil {
				log.Panic(err)
			}
			if !b {
				return false
			}
		case "pid":
			b, err := regexp.MatchString("^\\d{9}$", m[key])
			if err != nil {
				log.Panic(err)
			}
			if !b {
				return false
			}
		case "cid":
			continue
		}
	}
	return true
}
