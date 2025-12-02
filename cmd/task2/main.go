package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type idRange struct{ start, end int64 }

// This one is efficient and I'm happy with it.
func isSeqTwice(num int64) bool {
	// Log10(num) + 1 = the number of digits
	log := int64(math.Log10(float64(num)))
	nOfDigits := log + 1
	// odd number of digits -> def not a sequence repeated twice
	if nOfDigits%2 != 0 {
		return false
	}

	// don't like this name. When we divide input num by this,
	// we get a number twice as short (before decimal point)
	halfDigits := int64(math.Pow(10, float64(nOfDigits/2)))

	seq := num / halfDigits
	return seq*halfDigits+seq == num
}

// This meh works, nothing more to say.
func isSeqN(num int64) bool {
	numStr := strconv.FormatInt(num, 10)

	for i := 1; i <= len(numStr)/2; i++ {
		pattern := fmt.Sprintf("^(%s)+$", numStr[:i])

		if match, _ := regexp.MatchString(pattern, numStr); match {
			return true
		}
	}

	return false
}

// This is ~55~ 70 times faster than meh-version.
func isSeqNLogs(num int64) bool {
	log := int64(math.Log10(float64(num)))
	nOfDigits := log + 1 // number of digits in num

	// I will try with prefixes with length 1 (e.g. 4 for 444)
	// and up to half of nOfDigits, that's 3 for 123123.
	for prefixLen := int64(1); prefixLen <= nOfDigits/2; prefixLen++ {
		// even prefix length don't work for odd number lengths.
		if prefixLen%2 == 0 && nOfDigits%2 == 1 {
			continue
		}

		denom := int64(math.Pow(10, float64(nOfDigits-prefixLen)))
		prefixNum := num / denom
		multiplier := int64(math.Pow(10, float64(prefixLen)))

		rebuiltNum := int64(0)
		for j := int64(0); j < nOfDigits/prefixLen; j++ {
			rebuiltNum *= multiplier
			rebuiltNum += prefixNum
		}
		
		if rebuiltNum == num {
			return true
		}
	}

	return false
}
func parse(r io.Reader) ([]idRange, error) {
	reader := bufio.NewReader(r)
	result := []idRange{}
	for {
		str, err := reader.ReadString(',')
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return nil, err
			}
		}
		if len(str) == 0 {
			return result, nil
		}

		nums := strings.Split(str[:len(str)-1], "-")
		start, err := strconv.ParseInt(nums[0], 10, 64)
		if err != nil {
			return nil, err
		}

		end, err := strconv.ParseInt(nums[1], 10, 64)
		if err != nil {
			return nil, err
		}

		result = append(result, idRange{start, end})
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

	idRanges, err := parse(r)
	if err != nil {
		panic(err)
	}

	var answer int64
	for _, rng := range idRanges {
		for id := rng.start; id <= rng.end; id++ {
			// if isSeqTwice(id) {
			if isSeqNLogs(id) {
				answer += id
			}
		}
	}

	fmt.Println(answer)
}
