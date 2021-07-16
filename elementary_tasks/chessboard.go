package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const helloMsg = `you can type height and width of your chessboard
and a symbol to be used in it
`

func main() {
	fmt.Println(helloMsg)

	var err error
	consoleScan := func(txt string) int {
		if err != nil {
			return 0
		}
		num, e := scanInt(txt)
		err = e
		if num < 0 {
			err = errors.New("should be positive integers")
		}
		return num
	}

	i := consoleScan("input height: ")
	w := consoleScan("input width: ")
	if err != nil {
		fmt.Println("should be positive integers")
		return
	}

	s := scanString("symbol: ")
	row := createRow(w)
	printChessboard(
		formatOddRow(row, s),
		formatEvenRow(row, s),
		i)
}

//input: width
func createRow(w int) int {
	row := 1
	for i := 1; i < w*2; i++ {
		row = row << 1
		if i%2 == 0 {
			row++
		}
	}

	return row
}

func formatOddRow(row int, s string) string {
	r := "0" + strconv.FormatInt(int64(row>>1), 2)
	r = format(r, s)
	return r
}

func formatEvenRow(row int, s string) string {
	r := strconv.FormatInt(int64(row), 2)
	r = format(r, s)
	return r
}

// odd and even row, height
func printChessboard(odd, even string, h int) {
	for i := 1; i <= h; i++ {
		if i%2 == 0 {
			fmt.Println(odd)
			continue
		}
		fmt.Println(even)
	}
}

// input: question
func scanString(q string) string {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print(q)
	sc.Scan()
	str := sc.Text()
	return str
}

func scanInt(q string) (int, error) {
	str := scanString(q)
	res, err := strconv.ParseInt(str, 10, 32)
	return int(res), err
}

func format(str string, sym string) string {
	str = strings.ReplaceAll(str, "0", " ")
	str = strings.ReplaceAll(str, "1", sym)
	return str
}
