package string_sum

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	var summ = 0
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	arrOfDigits := re.FindAllString(input, -1)
	for _, elem := range arrOfDigits {
		value, _ := strconv.Atoi(elem)
		summ += value
	}
	return string(summ), nil
}
