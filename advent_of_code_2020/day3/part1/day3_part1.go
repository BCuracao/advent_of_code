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

	fmt.Println(len(ar))    // column
	fmt.Println(len(ar[0])) // row

	var dx, dy, t int = 3, 1, 0
	for y, x := 0, 0; y < len(ar); y, x = y+dy, x+dx {
		if string(ar[y][x%len(ar[y])]) == "#" {
			t++
		}
	}
	fmt.Println("Trees encountered: ", t)
}
