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

	horizonal, depth := 0, 0

	for sc.Scan() {
		s := strings.Split(sc.Text(), " ")
		switch s[0] {
		case "forward":
			h, _ := strconv.Atoi(s[1])
			horizonal += h

		case "down":
			d, _ := strconv.Atoi(s[1])
			depth += d

		case "up":
			u, _ := strconv.Atoi(s[1])
			depth -= u
		}
	}
	fmt.Println(horizonal * depth)
}
