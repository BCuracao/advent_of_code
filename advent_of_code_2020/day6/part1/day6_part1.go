package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var m map[rune]int
var mapp map[rune]int

func main() {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	groups := make([]string, 0)
	for _, s := range strings.Split(string(f), "\n\n") {
		groups = append(groups, s)
	}
	sum := countYesByGroups(groups)
	fmt.Println("Solution: ", sum)
}

func countYesByGroups(groups []string) int {
	g, sum := 0, 0

	for _, s := range groups {
		c, n := 0, 1
		m = make(map[rune]int, 26)
		for _, c := range s {
			if c == '\\' {
				n++
			}
			m[c]++
		}
		fmt.Println("backslash: ", n)
		for k, _ := range m {
			if m[k] > 1 {
				c++
			}
		}
		g++
		sum += c
		//fmt.Println("Group Number: ", g, "yes answers: ", c)
	}
	return sum
}
