package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type (
	direction int
	movement  struct {
		d direction
		c int
	}
)

const (
	left direction = iota
	right

	dialSize      = 100
	startPosition = 50
)

func (m movement) apply(position int) int {
	position += 100
	if m.d == left {
		position -= m.c
	} else {
		position += m.c
	}

	return position % 100
}
func (m movement) apply2(startPosition int) (position, zeros int) {
	zeros, c := m.c/dialSize, m.c%dialSize
	position = startPosition
	// No steps
	if c == 0 {
		return position, zeros
	}

	// We are guaranteed to make less that dialSize steps
	if m.d == left {
		position -= c
	} else {
		position += c
	}

	if position < 0 && startPosition == 0 {
		position += dialSize
	} else if position < 0 {
		position += dialSize
		zeros++
	} else if position >= dialSize {
		position -= dialSize
		zeros++
	} else if position == 0 {
		zeros++
	}

	return position, zeros
}

func solve(pos int, movs []movement) {
	zeroCount := 0
	for _, mov := range movs {
		pos = mov.apply(pos)
		if pos == 0 {
			zeroCount++
		}
	}
	fmt.Println(zeroCount)
}

func solve2(pos int, movs []movement) {
	zeroCount := 0
	for _, mov := range movs {
		var zeros int
		pos, zeros = mov.apply2(pos)
		zeroCount += zeros
	}
	fmt.Println(zeroCount)
}

func parse(r io.Reader) ([]movement, error) {
	res := make([]movement, 0, 32)
	reader := bufio.NewReader(r)
	for {
		line, ok, err := reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return res, nil
			}
			return nil, err
		} else if ok {
			return nil, errors.New("line too long")
		}
		if len(line) == 0 {
			continue
		}

		var d direction
		switch line[0] {
		case 'L':
			d = left
		case 'R':
			d = right
		default:
			return nil, errors.New("wrong direction")
		}

		steps, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			return nil, err
		}

		res = append(res, movement{d: d, c: steps})
	}

	return res, nil
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

	movs, err := parse(r)
	if err != nil {
		panic(err)
	}

	solve2(startPosition, movs)
}
