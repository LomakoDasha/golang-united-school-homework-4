package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	var summ = 0
	var emptyInputCustomError = fmt.Errorf(errorEmptyInput.Error())
	var wrongOperandsAmountCustomError = fmt.Errorf(errorNotTwoOperands.Error())
	cleanString := strings.TrimSpace(input)

	if len(cleanString) == 0 {
		return "", emptyInputCustomError
	}

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	arrOfDigits := re.FindAllString(cleanString, -1)

	if len(arrOfDigits) != 2 {
		return "", wrongOperandsAmountCustomError
	} else {
		for _, elem := range arrOfDigits {
			value, err := strconv.Atoi(elem)
			error := fmt.Errorf("Incorrect operand %w", err)
			if err != nil {
				return "", error
			}
			summ += value
		}
	}

	return strconv.Itoa(summ), nil
}
