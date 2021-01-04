// Package fb is the core logic of fizzbuzz algorithm
package fb

import (
	"errors"
	"strconv"
	"strings"
)

const start int64 = 1
const max int64 = 100
const base int = 10

// DoFizzBuzz returns all the number between 1 and limit.
//
// It replaces all multiples of int1 by str1, all multiples of int2 by str2.
// Multiples of both int1 and int2 are replaced by str1str2.
func DoFizzBuzz(int1, int2, limit int64, str1, str2 string) ([]string, error) {
	result := ""
	separator := ""

	if limit < start || limit > max {
		return nil, errors.New(
			"limit must be between" + strconv.FormatInt(start, base) + " and " + strconv.FormatInt(max, base))
	}

	for i := start; i <= limit; i++ {
		if i > start {
			separator = ","
		}
		if i%int1 == 0 && i%int2 == 0 {
			result += separator + str1 + str2
		} else if i%int1 == 0 {
			result += separator + str1
		} else if i%int2 == 0 {
			result += separator + str2
		} else {
			result += separator + strconv.FormatInt(i, base)
		}
	}

	return strings.Split(result, ","), nil
}
