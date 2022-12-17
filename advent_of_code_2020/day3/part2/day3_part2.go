package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	ar := make([]string, 0)
	for _, b := range bytes.Split(f, []byte("\n")) {
		ar = append(ar, string(b))
	}

	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	var dx, dy int
	sum := 1
	for i := 0; i < len(slopes); i++ {
		for j := 0; j < len(slopes[i])-1; j++ {
			dx, dy = slopes[i][j], slopes[i][j+1]
		}
		sum *= traverseSlope(dx, dy, ar)
	}
	fmt.Println("Product: ", sum)
}

func traverseSlope(dx, dy int, ar []string) int {

	t := 0
	for y, x := 0, 0; y < len(ar); y, x = y+dy, x+dx {
		if string(ar[y][x%len(ar[y])]) == "#" {
			t++
		}
	}
	return t
}
