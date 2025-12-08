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
		// used to be simply countRolls and no loop.
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

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if plan[i][j] == 0 {
				continue
			}

			incrementAround(neighbours, i, j)
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
			}
		}
	}

	return result
}

func incrementAroundRow(row []int64, i int, withCenter bool) {
	if i > 0 {
		row[i-1]++
	}

	if withCenter {
		row[i]++
	}

	if i < len(row)-1 {
		row[i+1]++
	}
}

func incrementAround(rows [][]int64, i, j int) {
	if i != 0 {
		incrementAroundRow(rows[i-1], j, true)
	}

	incrementAroundRow(rows[i], j, false)

	if i < len(rows[0])-1 {
		incrementAroundRow(rows[i+1], j, true)
	}
}
