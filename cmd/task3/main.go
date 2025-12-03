package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
)

func parseLine(line string) []int {
	res := make([]int, len(line))
	for i, c := range line {
		jolts := c - '0'
		res[i] = int(jolts)
	}
	return res
}

func parse(r io.Reader) ([][]int, error) {
	reader := bufio.NewReader(r)
	result := [][]int{}

	for {
		str, _, err := reader.ReadLine()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return nil, err
			}
		}
		if len(str) == 0 {
			return result, nil
		}
		bank := parseLine(string(str))
		result = append(result, bank)
	}
}

func main() {
	var r io.Reader = os.Stdin
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer file.Close()
		r = file
	}

	banks, err := parse(r)
	if err != nil {
		panic(err)
	}

	var answer int64
	for _, bank := range banks {
		answer += buildMax(bank, 12)
	}

	fmt.Println(answer)
}

func find2Batts(bank []int) int64 {
	maxPossible := make([]int, len(bank))
	maxSoFar := bank[0]
	// -1 because the last batt can't be used twice
	for i := 0; i < len(bank)-1; i++ {
		maxSoFar = max(bank[i], maxSoFar)
		maxPossible[i] = 10 * maxSoFar
	}

	// we iterate backwards. we have to look at "next" batteries only.
	maxSoFar = bank[len(bank)-1]
	for i := len(bank) - 1; i >= 0; i-- {
		maxPossible[i] += maxSoFar
		maxSoFar = max(bank[i], maxSoFar)
	}

	maxBattery := 0
	for _, v := range maxPossible {
		maxBattery = max(maxBattery, v)
	}
	return int64(maxBattery)
}

func buildMax(digits []int, nDigits int) int64 {
	if nDigits == 0 {
		panic("unreachable")
	}

	total := len(digits)
	maxIndex := total + 1 - nDigits

	maxDigit, maxDigitIndex := digits[0], 0
	for i := 0; i < maxIndex; i++ {
		if digits[i] > maxDigit {
			maxDigit, maxDigitIndex = digits[i], i
		}
	}

	if maxDigitIndex == total-1 {
		return int64(maxDigit)
	}

	if nDigits == 1 {
		return int64(maxDigit)
	}

	right := buildMax(digits[maxDigitIndex+1:], nDigits-1)
	res := int64(maxDigit)*int64(math.Pow(10, float64(nDigits-1))) + right
	return res
}
