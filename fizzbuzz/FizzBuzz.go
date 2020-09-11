package fizzbuzz

import (
	"strconv"
)

// FizzBuzz : fizzbuzz fucntion
func FizzBuzz(param int) string {
	returnStr := ""

	if param < 1 {
		returnStr = strconv.Itoa(param)
	} else if param%3 == 0 && param%5 == 0 {
		returnStr = "FizzBuzz"
	} else if param%5 == 0 {
		returnStr = "Buzz"
	} else if param%3 == 0 {
		returnStr = "Fizz"
	} else {
		returnStr = strconv.Itoa(param)
	}

	return returnStr
}
