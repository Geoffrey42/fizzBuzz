package fizzbuzz

import "strconv"

func doFizzBuzz(int1, int2, limit int, str1, str2 string) string {
	result := "1"

	for i := 2; i <= limit; i++ {
		if i%int1 == 0 && i%int2 == 0 {
			result += "," + str1 + str2
		} else if i%int1 == 0 {
			result += "," + str1
		} else if i%int2 == 0 {
			result += "," + str2
		} else {
			result += "," + strconv.Itoa(i)
		}
	}

	return result
}
