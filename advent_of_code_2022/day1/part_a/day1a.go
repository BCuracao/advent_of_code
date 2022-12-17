package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	in, _ := os.Open("../input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)

	var elves [][]string

	for sc.Scan() {
		num := strings.Split(sc.Text(), "")
		elves = append(elves, num)
	}
}
