package main

import "testing"

func TestFind2Batts(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{3, 1, 4, 2}, 42},
		{parseLine("987654321111111"), 98},
		{parseLine("811111111111119"), 89},
		{parseLine("234234234234278"), 78},
		{parseLine("818181911112111"), 92},
	}
	for _, tc := range testCases {
		got := find2Batts(tc.input)
		if got != tc.expected {
			t.Fail()
		}
	}
}

func TestBuildMax(t *testing.T) {
	testCases := []struct {
		input    []int
		nDigits  int
		expected int64
	}{
		{parseLine("987654321111111"), 2, 98},
		{parseLine("811111111111119"), 2, 89},
		{parseLine("234234234234278"), 2, 78},
		{parseLine("818181911112111"), 2, 92},
		{parseLine("987654321111111"), 12, 987654321111},
		{parseLine("811111111111119"), 12, 811111111119},
		{parseLine("234234234234278"), 12, 434234234278},
		{parseLine("818181911112111"), 12, 888911112111},
	}

	for _, tc := range testCases {
		got := buildMax(tc.input, tc.nDigits)
		if got != tc.expected {
			buildMax(tc.input, tc.nDigits)
			t.Fail()
		}
	}
}
