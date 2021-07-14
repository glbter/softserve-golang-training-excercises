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
	consoleScanner := bufio.NewScanner(os.Stdin)
	consoleScanner.Scan()
	path := consoleScanner.Text()

	file, err := os.Open(strings.TrimSpace(path))
	if err != nil {
		fmt.Println("could not open the file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	algoName := scanner.Text()

	algo, err := getAlgorithm(algoName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var count int
	for scanner.Scan() {
		data, err := scan(scanner)
		if err != nil || !validTicket(data) {
			continue
		}

		isLucky := algo(splitToDigits(data))
		if isLucky {
			count++
		}
	}

	fmt.Println("--Result--")
	fmt.Println(fmt.Sprint("Amount of lucky tickets: ", count))
}

func getAlgorithm(algoName string) (func([]int) bool, error) {
	switch strings.ToLower(algoName) {
	case "piter":
		return validateHardFormula, nil
	case "moskow":
		return validateEasyFormula, nil
	default:
		return nil, errors.New("not valid algorithm")
	}
}

func validateEasyFormula(ticket []int) bool {
	var left_sum, right_sum int
	for i, elem := range ticket {
		if i < 3 {
			left_sum += elem
		} else {
			right_sum += elem
		}
	}

	return left_sum == right_sum
}

func validateHardFormula(ticket []int) bool {
	var even_sum, odd_sum int
	for i, elem := range ticket {
		if i%2 == 0 {
			even_sum += elem
		} else {
			odd_sum += elem
		}
	}

	return even_sum == odd_sum
}

func reverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func splitToDigits(n int) []int {
	var ret []int

	for n != 0 {
		ret = append(ret, n%10)
		n /= 10
	}

	reverseInt(ret)

	return ret
}

func scan(scanner *bufio.Scanner) (int, error) {
	input := scanner.Text()
	result, err := strconv.Atoi(input)
	return int(result), err
}

func validTicket(num int) bool {
	return 1e5 <= num && num <= 1e6
}
