package task6

import (
	"errors"
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

func ValidTicket(num int) bool {
	return 1e5 <= num && num <= 1e6
}
