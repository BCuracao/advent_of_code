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
	firstValue, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	secondValue, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	thirdValue, _ := strconv.Atoi(sc.Text())

	//     A    B    C
	// 1: 191
	// 2: 185	185
	// 3: 188	188  188
	// 4:    	189  189  189
	// 5:	     	 204  204
	// 6:                 213

	i, increases := 0, 0
	for sc.Scan() {
		newthirdValue, _ := strconv.Atoi(sc.Text())

		if newthirdValue > firstValue {
			increases++
		}
		fmt.Println("first: ", firstValue, " second: ", secondValue, " third: ", newthirdValue)
		firstValue = secondValue
		secondValue = thirdValue
		thirdValue = newthirdValue
		i++
	}
	fmt.Println(increases)
}
