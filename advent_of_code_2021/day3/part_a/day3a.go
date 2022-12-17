package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.Open("../input.txt")
	defer in.Close()
	sc := bufio.NewScanner(in)

	var bits [][]string

	for sc.Scan() {
		num := strings.Split(sc.Text(), "")
		bits = append(bits, num)
	}
	cOnes := searchOneBits(bits)

	var numOfOnes bytes.Buffer
	for i := 0; i < len(cOnes); i++ {
		if cOnes[i] > 500 {
			numOfOnes.WriteString("1")
		} else {
			numOfOnes.WriteString("0")
		}
	}

	gammaRate, _ := strconv.ParseInt(numOfOnes.String(), 2, 32)

	var epsilonBuffer bytes.Buffer
	for _, digit := range numOfOnes.String() {
		epsilonBuffer.WriteString(string(1 ^ digit))
	}
	epsilonRate, _ := strconv.ParseInt(epsilonBuffer.String(), 2, 32)
	fmt.Println("The multiplied result: ", gammaRate*epsilonRate)
}

func searchOneBits(bits [][]string) []int {
	sumOfOnes := make([]int, 12)
	for i := 0; i < len(bits); i++ {
		for j := 0; j < len(bits[0]); j++ {
			if bits[i][j] == "1" {
				sumOfOnes[j] += 1
			}
		}
	}
	return sumOfOnes
}
