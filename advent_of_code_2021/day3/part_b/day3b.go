package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("../input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)

	data := []string{}

	for sc.Scan() {
		data = append(data, sc.Text())
	}
	//fmt.Println(data)
	bitCriteria(data)
}

func bitCriteria(data []string) {
	zeroes, ones, k := []string{}, []string{}, 0
	for i := 0; i < len(data); i++ {
		runes := []rune(data[i])
		for j := k; j < len(runes); {
			if runes[k] == '0' {
				zeroes = append(zeroes, data[k])
			} else {
				ones = append(ones, data[k])
			}
		}
	}
	fmt.Println("THE ZEROES:")
	fmt.Println("")
	fmt.Println(zeroes)
}
