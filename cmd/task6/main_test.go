package main

import (
	"strings"
	"testing"
)

const raw = "" +
	"12  3 4 \n" +
	"56 7   8\n" +
	"*  +  * \n"

func TestParse(t *testing.T) {
	expected := [][]int64{
		[]int64{26, 15},
		[]int64{3, 7},
		[]int64{8, 4},
	}
	parsed, err := parse(strings.NewReader(raw))
	if err != nil {
		t.Fail()
	}

	for i, row := range expected {
		for j, expectedNum := range row {
			if expectedNum != parsed.num2[i][j] {
				t.Fail()
			}
		}
	}

}
