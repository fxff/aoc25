package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func parseLine(line []byte) []int64 {
	res := make([]int64, len(line))
	for i, c := range line {
		if c == '@' {
			res[i] = 1
		}
	}
	return res
}

func parse(r io.Reader) ([][]int64, error) {
	reader := bufio.NewReader(r)
	result := [][]int64{}

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
		line := parseLine(str)
		result = append(result, line)
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

	plan, err := parse(r)
	if err != nil {
		panic(err)
	}

	var totalRolls int64
	for {
		neighbours := buildNeighbours(plan)
		removed := removeRolls(plan, neighbours)
		if removed == 0 {
			break
		}
		totalRolls += removed
	}

	fmt.Println(totalRolls)
}

func buildNeighbours(plan [][]int64) [][]int64 {
	h, w := len(plan), len(plan[0])
	neighbours := make([][]int64, h, h)
	for i := range neighbours {
		neighbours[i] = make([]int64, w)
	}

	for i, row := range plan {
		for j, v := range row {
			if v == 0 {
				continue
			}
			if i != 0 {
				incrementNeighbourCount(neighbours[i-1], j)
			}
			neighbours[i][j]--
			incrementNeighbourCount(neighbours[i], j)
			if i != len(neighbours)-1 {
				incrementNeighbourCount(neighbours[i+1], j)
			}
		}
	}
	return neighbours
}

func removeRolls(plan, neighbours [][]int64) int64 {
	var result int64
	for i, row := range plan {
		for j, nRolls := range row {
			if nRolls > 0 && neighbours[i][j] < 4 {
				row[j]--
				result++
			}
		}
	}

	return result
}

func countRolls(plan, neighbours [][]int64) int64 {
	var result int64
	for i, row := range plan {
		for j, n := range row {
			if n != 0 && neighbours[i][j] < 4 {
				result++
				println("less4: ", i, j)
			}
		}
	}
	return result
}
func incrementNeighbourCount(row []int64, pos int) {
	if pos > 0 {
		row[pos-1]++
	}
	row[pos]++
	if pos < len(row)-1 {
		row[pos+1]++
	}
}
