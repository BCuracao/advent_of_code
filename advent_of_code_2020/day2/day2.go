package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	res := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		//fmt.Println(sc.Text())
		str := sc.Text()
		cLeftIdx := strings.Index(str, "-")
		cRightIdx := strings.Index(str, " ")
		lIdx := strings.Index(str, ":")

		min, _ := strconv.Atoi(str[:cLeftIdx])
		max, _ := strconv.Atoi(str[cLeftIdx+1 : cRightIdx])
		letter := str[cRightIdx+1 : lIdx]
		pSeries := str[strings.LastIndex(str, " ")+1:]

		/* Part 1:
		c := 0
		for _, l := range pSeries {
			r := rune(letter[0])
			if l == r {
				c++
			}
		}
		if c >= min && c <= max {
			res++
		}
		*/

		// Part 2:
		fmt.Println("index min: ", min, "max :", max, " Sletter: ", letter, " - ", pSeries, " min :", string(pSeries[min-1]), " max: ", string(pSeries[max-1]))
		if string(pSeries[min-1]) == letter &&
			string(pSeries[max-1]) != letter {
			res++
		} else if string(pSeries[max-1]) == letter &&
			string(pSeries[min-1]) != letter {
			res++
		}
	}
	fmt.Println("The result is: ", res)
}
