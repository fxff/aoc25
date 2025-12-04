package main

import (
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	const ex = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`
	r := strings.NewReader(ex)
	plan, _ := parse(r)
	var totalRolls int64
	for {
		neighbours := buildNeighbours(plan)
		removed := removeRolls(plan, neighbours)
		if removed == 0 {
			break
		}
		totalRolls += removed
	}

	if totalRolls != 43 {
		t.Fail()
	}
}
