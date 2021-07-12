package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var err error
	consoleScan := func(text string) int {
		if err != nil {
			return 0
		}
		num, e := scan(text)
		err = e
		if num < 0 {
			err = errors.New("should be positive integers")
		}
		return num
	}

	height := consoleScan("input height: ")
	width := consoleScan("input width: ")
	if err != nil {
		fmt.Println("should be positive integers")
		return
	}

	symbol := scanString("symbol: ")
	row := createRow(width)
	printChessboard(
		formatOddRow(row, symbol),
		formatEvenRow(row, symbol),
		height)
}

func createRow(width int) int {
	row := 1
	for w := 1; w < width*2; w++ {
		row = row << 1
		if w%2 == 0 {
			row++
		}
	}

	return row
}

func formatOddRow(row int, symbol string) string {
	odd_row := "0" + strconv.FormatInt(int64(row>>1), 2)
	odd_row = format(odd_row, symbol)
	return odd_row
}

func formatEvenRow(row int, symbol string) string {
	even_row := strconv.FormatInt(int64(row), 2)
	even_row = format(even_row, symbol)
	return even_row
}

func printChessboard(odd_row, even_row string, height int) {
	for h := 1; h <= height; h++ {
		if h%2 == 0 {
			fmt.Println(odd_row)
		} else {
			fmt.Println(even_row)
		}
	}
}

func scanString(question string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(question)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func scan(question string) (int, error) {
	input := scanString(question)
	result, err := strconv.ParseInt(input, 10, 32)
	return int(result), err
}

func format(str string, symbol string) string {
	str = strings.ReplaceAll(str, "0", " ")
	str = strings.ReplaceAll(str, "1", symbol)
	return str
}
