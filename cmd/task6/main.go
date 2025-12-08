package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/fxff/aoc25/internal/app"
)

type (
	action int
	input  struct {
		nums    [][]int64
		num2    [][]int64
		actions []action
	}
)

const (
	actionAdd action = iota
	actionMul
)

func parse(r io.Reader) (*input, error) {
	var result = &input{}
	var lines []string

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// part 1
	for _, line := range lines {
		justStarted := true
		lineScanner := bufio.NewScanner(strings.NewReader(line))
		lineScanner.Split(bufio.ScanWords)

		for lineScanner.Scan() {
			token := lineScanner.Text()
			if justStarted && token != "+" && token != "*" {
				justStarted = false
				result.nums = append(result.nums, make([]int64, 0, 1024))
			}

			switch token {
			case "+":
				result.actions = append(result.actions, actionAdd)
			case "*":
				result.actions = append(result.actions, actionMul)
			default:
				num, err := strconv.ParseInt(token, 10, 64)
				if err != nil {
					return nil, err
				}

				result.nums[len(result.nums)-1] = append(result.nums[len(result.nums)-1], num)
			}
		}
		if err := lineScanner.Err(); err != nil {
			return nil, err
		}
	}

	// part 2
	result.num2 = make([][]int64, len(result.actions))
	groupPointer := len(result.actions) - 1
	for i := len(lines[0]) - 1; i >= 0; i-- {
		number := int64(0)
		isSeparator := true

		for j := 0; j < len(lines)-1; j++ {
			chr, num := lines[j][i], int64(lines[j][i]-'0')
			if chr == ' ' {
				continue
			}
			isSeparator = false
			number *= 10
			number += num
		}

		if isSeparator {
			groupPointer--
			continue
		}

		result.num2[groupPointer] = append(result.num2[groupPointer], number)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func solve(input *input) int64 {
	var result int64
	for i, action := range input.actions {
		var intermediate int64

		if action == actionMul {
			intermediate = 1
		}

		for _, row := range input.nums {
			num := row[i]

			switch action {
			case actionMul:
				intermediate *= num
			case actionAdd:
				intermediate += num
			}
		}

		result += intermediate
	}

	return result
}
func solve2(input *input) int64 {
	var result int64
	for i, action := range input.actions {
		var intermediate int64

		if action == actionMul {
			intermediate = 1
		}

		for _, num := range input.num2[i] {
			switch action {
			case actionMul:
				intermediate *= num
			case actionAdd:
				intermediate += num
			}
		}

		result += intermediate
	}

	return result
}

func main() {
	app.New(parse, solve, solve2).Run()
}
