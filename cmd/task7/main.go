package main

import (
	"bufio"
	"io"

	"github.com/fxff/aoc25/internal/app"
)

type (
	cell  int
	input struct {
		cells [][]cell
	}
)

func (c cell) String() string {
	switch c {
	case empty:
		return "."
	case emitter:
		return "s"
	case splitter:
		return "^"
	case beam:
		return "|"
	default:
		panic("panik")
	}
}

const (
	empty cell = iota
	emitter
	splitter
	beam
)

func parse(r io.Reader) (*input, error) {
	var result = &input{}

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		cellRow := make([]cell, len(line))
		for i, chr := range line {
			switch chr {
			case '.':
				cellRow[i] = empty
			case 'S':
				cellRow[i] = emitter
			case '^':
				cellRow[i] = splitter
			}
		}

		result.cells = append(result.cells, cellRow)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func solve(input *input) int64 {
	w := len(input.cells[0])
	state := make([]cell, w)

	result := int64(0)
	for _, row := range input.cells {
		for i, cell := range row {
			switch cell {
			case empty:
				continue
			case emitter:
				state[i] = beam
			case splitter:
				if state[i] != beam {
					continue
				}

				result++
				state[i] = empty
				if i > 0 {
					state[i-1] = beam
				}
				if i < w-1 {
					state[i+1] = beam
				}
			}
		}
	}

	return result
}
func solve2(input *input) int64 {
	w := len(input.cells[0])
	// past represents timelines that could have
	// led to this coordinate
	past := make([]int64, w)

	result := int64(0)
	for _, row := range input.cells {
		for i, cell := range row {
			switch cell {
			case empty:
				continue
			case emitter:
				past[i] = 1
			case splitter:
				// tbh we can omit size check here,
				// ain't that sweet

				if i > 0 {
					past[i-1] += past[i]
				}
				if i < w-1 {
					past[i+1] += past[i]
				}
				past[i] = 0
			}
		}
	}

	for _, v := range past {
		result += v
	}
	return result
}

func main() {
	app.New(parse, solve, solve2).Run()
}
