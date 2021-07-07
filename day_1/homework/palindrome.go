package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	str, err := scan("Input string: ")
	if err != nil {
		fmt.Println("not a numerical string")
		fmt.Println("0")
		return
	}

	if len(str) < 2 {
		fmt.Println("too short string")
		fmt.Println("0")
		return
	}

	fmt.Println("easy level")
	singleCount := isSinglePalindrome(str)
	if singleCount == 0 {
		fmt.Println(singleCount)
	}

	fmt.Println("hard level")
	length := len(str)

	mem := make([][]int, length)
	for i := range mem {
		mem[i] = make([]int, length)
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			mem[i][j] = -1
		}
	}

	var count int
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			isPal := isPalindrome(str, mem, i, j)
			count += isPal
			if isPal == 1 {
				fmt.Println(str[i : j+1])
			}
		}
	}

	if count == 0 {
		fmt.Println(count)
	}
}

func isPalindrome(str string, mem [][]int, i int, j int) int {
	// base condition
	if i > j {
		return 1
	}

	// if counted
	if mem[i][j] != -1 {
		return mem[i][j]
	}

	// first and last not equal
	if str[i] != str[j] {
		mem[i][j] = 0
		return 0
	}

	mem[i][j] = isPalindrome(str, mem, i+1, j-1)
	return mem[i][j]
}

func reverseStr(str string) string {
	s := strings.Split(str, "")

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return strings.Join(s, "")
}

func isSinglePalindrome(str string) int {
	if str == reverseStr(str) {
		fmt.Println(str)
		return 1
	}

	return 0
}

func scan(question string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(question)
	scanner.Scan()
	input := scanner.Text()
	_, err := strconv.Atoi(input)
	return input, err
}
