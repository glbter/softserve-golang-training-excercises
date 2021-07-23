package scan

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrNotPosInt      = errors.New("should be positive integer")
	ErrNotPosNum      = errors.New("should be positive number")
	ErrScannerTrouble = errors.New("scanner returned false")
)

// input: question

func ScanString(sc *bufio.Scanner, q string) (string, bool) {
	fmt.Print(q)
	ok := sc.Scan()
	return sc.Text(), ok
}

func ScanPositiveInt(sc *bufio.Scanner, q string) (int, error) {
	str, ok := ScanString(sc, q)
	if !ok {
		return 0, ErrScannerTrouble
	}
	res, err := strconv.ParseInt(str, 10, 32)
	if res <= 0 {
		err = ErrNotPosInt
	}

	return int(res), err
}

func ScanPositiveFloat(sc *bufio.Scanner, q string) (float32, error) {
	data, ok := ScanString(sc, q)
	if !ok {
		return 0, ErrScannerTrouble
	}
	res, err := strconv.ParseFloat((data), 32)
	if res <= 0 {
		err = ErrNotPosNum
	}
	return float32(res), err
}

func AskContinue(sc *bufio.Scanner) bool {
	data, _ := ScanString(sc, "Continue? ")
	cont := strings.ToLower(strings.TrimSpace(data))
	return cont == "yes" || cont == "y"
}
