package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	height, err := scan("input height:")
	if err != nil {
		fmt.Println("not a number")
		return
	}

	width, err := scan("input width:")
	if err != nil {
		fmt.Println("not a number")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("input symbol:")
	scanner.Scan()
	symbol := scanner.Text()

	row := 1
	for w := 1; w < width*2; w++ {
		row = row << 1
		if w%2 == 0 {
			row++
		}
	}

	odd_row := "0" + strconv.FormatInt(int64(row>>1), 2)
	odd_row = format(odd_row, symbol)

	even_row := strconv.FormatInt(int64(row), 2)
	even_row = format(even_row, symbol)

	for h := 1; h <= height; h++ {
		if h%2 == 0 {
			fmt.Println(odd_row)
		} else {
			fmt.Println(even_row)
		}
	}
}

func scan(question string) (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(question)
	scanner.Scan()
	input := scanner.Text()
	result, err := strconv.ParseInt(input, 10, 32)
	return int(result), err
}

func format(str string, symbol string) string {
	str = strings.ReplaceAll(str, "0", " ")
	str = strings.ReplaceAll(str, "1", symbol)
	return str
}
