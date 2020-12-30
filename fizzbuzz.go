package fizzbuzz

import (
	"errors"
	"strconv"
)

const start = 1
const max = 100

func doFizzBuzz(int1, int2, limit int, str1, str2 string) (string, error) {
	result := strconv.Itoa(start)
	separator := ","

	if limit < start || limit > max {
		return "", errors.New(
			"limit must be between" + result + " and " + strconv.Itoa(max))
	}

	for i := 2; i <= limit; i++ {
		if i%int1 == 0 && i%int2 == 0 {
			result += separator + str1 + str2
		} else if i%int1 == 0 {
			result += separator + str1
		} else if i%int2 == 0 {
			result += separator + str2
		} else {
			result += separator + strconv.Itoa(i)
		}
	}

	return result, nil
}
