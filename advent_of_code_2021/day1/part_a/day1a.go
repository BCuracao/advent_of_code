package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in, _ := os.Open("../input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)

	sc.Scan()
	prev, _ := strconv.Atoi(sc.Text())

	res := 0
	for sc.Scan() != false {
		n, _ := strconv.Atoi(sc.Text())
		if n > prev {
			res++
		}
		prev = n
	}
	fmt.Println("depth increase ", res, " times")
}
