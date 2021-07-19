package scan

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
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
		err = errors.New("should be positive integers")
	}
	return int(res), err
}

func ScanPositiveFloat(sc *bufio.Scanner, q string) (float32, error) {
	data := ScanString(sc, q)
	res, err := strconv.ParseFloat((data), 32)
	if res <= 0 {
		err = errors.New("negative lenght")
	}
	return float32(res), err
}

func AskContinue(sc *bufio.Scanner) bool {
	data := ScanString(sc, "Continue? ")
	cont := strings.ToLower(strings.TrimSpace(data))
	return cont == "yes" || cont == "y"
}
