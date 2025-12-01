package main

import "testing"

func TestApply(t *testing.T) {
	testCases := []struct {
		start int
		mov   movement
		end   int
	}{
		{start: 99, mov: movement{d: left, c: 1}, end: 98},
		{start: 99, mov: movement{d: right, c: 1}, end: 0},
		{start: 0, mov: movement{d: right, c: 1}, end: 1},
	}

	for _, tc := range testCases {
		got := tc.mov.apply(tc.start)
		if got != tc.end {
			t.Fatal()
		}
	}
}
func TestApply2(t *testing.T) {
	testCases := []struct {
		start      int
		mov        movement
		end, zeros int
	}{
		{start: 99, mov: movement{d: left, c: 1}, end: 98},
		{start: 0, mov: movement{d: right, c: 1}, end: 1},
		{start: 0, mov: movement{d: left, c: 1}, end: 99},
		{start: 1, mov: movement{d: left, c: 1}, end: 0, zeros: 1},
		{start: 99, mov: movement{d: right, c: 1}, end: 0, zeros: 1},
		{start: 0, mov: movement{d: right, c: 100}, end: 0, zeros: 1},
		{start: 1, mov: movement{d: right, c: 100}, end: 1, zeros: 1},
		{start: 1, mov: movement{d: right, c: 200}, end: 1, zeros: 2},
		{start: 1, mov: movement{d: left, c: 200}, end: 1, zeros: 2},
		{start: 99, mov: movement{d: right, c: 201}, end: 0, zeros: 3},
		//{start: 1, mov: movement{d: left, c: 200}, end: 1, zeros: 2},
	}

	for _, tc := range testCases {
		gotEnd, gotZeros := tc.mov.apply2(tc.start)

		if gotEnd != tc.end || gotZeros != tc.zeros {
			tc.mov.apply2(tc.start)
			t.Fatal()
		}
	}
}
