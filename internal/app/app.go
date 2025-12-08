package app

import (
	"fmt"
	"io"
	"os"
)

type (
	ParseFunc[T any] func(io.Reader) (T, error)
	SolveFunc[T any] func(T) int64
	App[T any]       struct {
		parse  ParseFunc[T]
		solve  SolveFunc[T]
		solve2 SolveFunc[T]
	}
)

func New[T any](
	parse ParseFunc[T],
	solve SolveFunc[T],
	solve2 SolveFunc[T],
) *App[T] {
	return &App[T]{
		parse:  parse,
		solve:  solve,
		solve2: solve2,
	}
}

func (a *App[T]) Run() {
	var r io.Reader = os.Stdin
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer file.Close()
		r = file
	}

	input, err := a.parse(r)
	if err != nil {
		panic(err)
	}

	answer := a.solve(input)
	fmt.Printf("1: %d\n", answer)

	answer = a.solve2(input)
	fmt.Printf("2: %d\n", answer)
}
