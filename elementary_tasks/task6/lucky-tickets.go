package task6

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

func GetAlgorithm(algo string) (func([]int) bool, error) {
	switch strings.ToLower(algo) {
	case "piter":
		return ValidateHardFormula, nil
	case "moskow":
		return ValidateEasyFormula, nil
	default:
		return nil, errors.New("not valid algorithm")
	}
}

// input ticket
func ValidateEasyFormula(t []int) bool {
	var lSum, rSum int
	for i, elem := range t {
		if i < 3 {
			lSum += elem
		} else {
			rSum += elem
		}
	}

	return lSum == rSum
}

// input ticket
func ValidateHardFormula(t []int) bool {
	var eSum, oSum int // even and odd
	for i, elem := range t {
		if i%2 == 0 {
			eSum += elem
		} else {
			oSum += elem
		}
	}

	return eSum == oSum
}

func reverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func SplitToDigits(n int) []int {
	var ret []int

	for n != 0 {
		ret = append(ret, n%10)
		n /= 10
	}

	reverseInt(ret)

	return ret
}

func Scan(sc *bufio.Scanner) (int, error) {
	str := sc.Text()
	got, err := strconv.Atoi(str)
	return int(got), err
}

func ValidTicket(num int) bool {
	return 1e5 <= num && num <= 1e6
}
