package main

import (
	"testing"
)

func TestSeqTwice(t *testing.T) {
	testCases := []struct {
		input    int64
		expected bool
	}{
		{1, false},
		{11, true},
		{99, true},
		{404, false},
		{1212, true},
		{1234512345, true},
	}

	for _, tc := range testCases {
		got := isSeqTwice(tc.input)
		if got != tc.expected {
			println(tc.input, got)
			t.Fatal()
		}
	}
}
func TestSeqN(t *testing.T) {
	testCases := []struct {
		input    int64
		expected bool
	}{
		{1, false},
		{11, true},
		{99, true},
		{404, false},
		{1212, true},
		{1234512345, true},
		{11, true},
		{111, true},
		{121212, true},
	}

	for _, tc := range testCases {
		got := isSeqN(tc.input)
		if got != tc.expected {
			t.Fatal()
		}
		got = isSeqNLogs(tc.input)
		if got != tc.expected {
			t.Fatal()
		}
	}
}

func BenchmarkSeqN(b *testing.B) {
	for range b.N {
		isSeqN(712712712)
	}
}

func BenchmarkSeqNLog(b *testing.B) {
	for range b.N {
		isSeqNLogs(712712712)
	}
}
