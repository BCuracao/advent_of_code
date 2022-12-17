package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

const YEAR = 2020

func main() {
	m := make(map[int]int)
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range bytes.Split(f, []byte("\n")) {
		n, err := strconv.Atoi(string(l))
		if err != nil {
			log.Println(err)
			continue
		}
		m[n] = 1
	}
	for y, _ := range m {
		fmt.Println(y)
		x := YEAR - y
		if _, ok := m[x]; ok {
			log.Printf("%d+%d=2020 / %d*%d=%d", x, y, x, y, x*y)
			fmt.Println()
		}
	}
}
