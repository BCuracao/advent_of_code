package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const ROWS = 128
const COLS = 8

var hId int = 0

func main() {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	x := 0
	arrRow := make([]int, 128)
	for i := range arrRow {
		arrRow[i] = x
		x++
	}
	y := 0
	arrCol := make([]int, 8)
	for i := range arrCol {
		arrCol[i] = y
		y++
	}
	//fmt.Println(arr)
	for _, s := range strings.Split(string(f), "\n") {
		r, c := splitRowCol(s)
		row := findRow(arrRow, r)
		col := findCol(arrCol, c)
		id := row*8 + col
		if id > hId {
			hId = id
		}
	}
	fmt.Println("Highest ID: ", hId)
}

func findRow(arr []int, s string) int {
	tmp := arr
	for _, r := range s {
		ch := string(r)
		if ch == "F" {
			tmp = tmp[:len(tmp)/2]
		}
		if ch == "B" {
			tmp = tmp[len(tmp)/2:]
		}
	}
	return int(tmp[0])
}

func findCol(arr []int, s string) int {
	tmp := arr
	for _, r := range s {
		ch := string(r)
		if ch == "L" {
			tmp = tmp[:len(tmp)/2]
		}
		if ch == "R" {
			tmp = tmp[len(tmp)/2:]
		}
	}
	return int(tmp[0])
}

func splitRowCol(s string) (string, string) {
	row := s[:len(s)-3]
	col := s[len(s)-3:]
	return row, col
}
