package scan

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrNotPosInt = errors.New("should be positive integer")
	ErrNotPosNum = errors.New("should be positive number")
)

// input: question
func ScanString(sc *bufio.Scanner, q string) string {
	fmt.Print(q)
	sc.Scan()
	return sc.Text()
}

func ScanPositiveInt(sc *bufio.Scanner, q string) (int, error) {
	str := ScanString(sc, q)
	res, err := strconv.ParseInt(str, 10, 32)
	if res <= 0 {
		err = ErrNotPosInt
	}
	return int(res), err
}

func ScanPositiveFloat(sc *bufio.Scanner, q string) (float32, error) {
	data := ScanString(sc, q)
	res, err := strconv.ParseFloat((data), 32)
	if res <= 0 {
		err = ErrNotPosNum
	}
	return float32(res), err
}

func AskContinue(sc *bufio.Scanner) bool {
	data := ScanString(sc, "Continue? ")
	cont := strings.ToLower(strings.TrimSpace(data))
	return cont == "yes" || cont == "y"
}
