package task6

import (
	"bufio"
	"errors"
	"strings"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/console/scan"
	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/integer"
	f "github.com/glbter/softserve-golang-training-excercises/elementary_tasks/task6/formula"
)

func GetAlgorithm(algo string) (FormulaValidator, error) {
	switch strings.ToLower(algo) {
	case "piter":
		return &f.HardFormula{}, nil
	case "moskow":
		return &f.EasyFormula{}, nil
	default:
		return nil, errors.New("not valid algorithm")
	}
}

type FormulaValidator interface {
	Validate(t []int) bool
}

func ValidTicket(num int) bool {
	return 1e5 <= num && num <= 1e6
}

func CountTickets(sc *bufio.Scanner, algo FormulaValidator) int {
	sc.Split(bufio.ScanWords)

	var c int
	for {
		num, err := scan.ScanPositiveInt(sc, "")
		if err == scan.ErrScannerTrouble {
			break
		}

		if err != nil || !ValidTicket(num) {
			continue
		}

		lucky := algo.Validate(integer.SplitToDigits(num))
		if lucky {
			c++
		}
	}

	return c
}
