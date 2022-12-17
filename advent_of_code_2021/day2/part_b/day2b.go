package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.Open("../input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)

	horizonal, depth, aim := 0, 0, 0

	for sc.Scan() {
		s := strings.Split(sc.Text(), " ")
		switch s[0] {
		case "forward":
			h, _ := strconv.Atoi(s[1])
			horizonal += h
			depth += aim * h

		case "down":
			a, _ := strconv.Atoi(s[1])
			aim += a

		case "up":
			a, _ := strconv.Atoi(s[1])
			aim -= a
		}
	}
	fmt.Println(horizonal * depth)
}
