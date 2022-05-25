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
	var emptyInputCustomError = fmt.Errorf("Error of empty input: %w", errorEmptyInput)
	var wrongOperandsAmountCustomError = fmt.Errorf("Error of more or less than two operands: %w", errorNotTwoOperands)
	cleanString := strings.TrimSpace(input)

	if len(cleanString) == 0 {
		return "", emptyInputCustomError
	}

	charactersRegexp := regexp.MustCompile(`[a-zA-Zа-яА-Я]`)
	arrOfCharacters := charactersRegexp.FindAllString(cleanString, -1)
	if len(arrOfCharacters) != 0 {
		_, err := strconv.Atoi(arrOfCharacters[0])
		letterExistError := fmt.Errorf("Letter exist in the string %w", err)
		if err != nil {
			return "", letterExistError
		}
	}

	digitsRegexp := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	arrOfDigits := digitsRegexp.FindAllString(cleanString, -1)

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
