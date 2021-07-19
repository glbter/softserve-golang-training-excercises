package task1

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//input: width
func CreateRow(w int) int {
	row := 1
	for i := 1; i < w*2; i++ {
		row = row << 1
		if i%2 == 0 {
			row++
		}
	}

	return row
}

func FormatEvenRow(row int, s string) string {
	r := "0" + strconv.FormatInt(int64(row>>1), 2)
	r = Format(r, s)
	return r
}

func FormatOddRow(row int, s string) string {
	r := strconv.FormatInt(int64(row), 2)
	r = Format(r, s)
	return r
}

// odd and even row, height
func MakeChessboard(odd, even string, h int) string {
	var b strings.Builder
	for i := 1; i <= h; i++ {
		if i != 1 {
			b.WriteString("\n")
		}
		if i%2 == 0 {
			b.WriteString(even)
			continue
		}
		b.WriteString(odd)
	}
	return b.String()
}

// input: question
func ScanString(sc *bufio.Scanner, q string) string {
	fmt.Print(q)
	sc.Scan()
	str := sc.Text()
	return str
}

func ScanPositiveInt(sc *bufio.Scanner, q string) (int, error) {
	str := ScanString(sc, q)
	res, err := strconv.ParseInt(str, 10, 32)
	if err == nil && res <= 0 {
		err = errors.New("should be positive integers")
	}
	return int(res), err
}

func Format(str string, sym string) string {
	str = strings.ReplaceAll(str, "0", " ")
	str = strings.ReplaceAll(str, "1", sym)
	return str
}
