package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const helloMsg = `program takes a file path as input
in the first line in file should be defined algorithm to count lucky tickets: moskow or piter
than line by line defined six number tickets`

func main() {
	fmt.Println(helloMsg)

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	path := sc.Text()

	file, err := os.Open(strings.TrimSpace(path))
	if err != nil {
		fmt.Println("could not open the file")
		return
	}
	defer file.Close()

	sc = bufio.NewScanner(file)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	algoN := sc.Text()

	algo, err := getAlgorithm(algoN)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var count int
	for sc.Scan() {
		data, err := scan(sc)
		if err != nil || !validTicket(data) {
			continue
		}

		lucky := algo(splitToDigits(data))
		if lucky {
			count++
		}
	}

	fmt.Println("--Result--")
	fmt.Println(fmt.Sprint("Amount of lucky tickets: ", count))
}

func getAlgorithm(algo string) (func([]int) bool, error) {
	switch strings.ToLower(algo) {
	case "piter":
		return validateHardFormula, nil
	case "moskow":
		return validateEasyFormula, nil
	default:
		return nil, errors.New("not valid algorithm")
	}
}

// input ticket
func validateEasyFormula(t []int) bool {
	var lSum, rSum int
	for i, elem := range t {
		if i < 3 {
			lSum += elem
		} else {
			rSum += elem
		}
	}

	return lSum == rSum
}

// input ticket
func validateHardFormula(t []int) bool {
	var eSum, oSum int // even and odd
	for i, elem := range t {
		if i%2 == 0 {
			eSum += elem
		} else {
			oSum += elem
		}
	}

	return eSum == oSum
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

func scan(sc *bufio.Scanner) (int, error) {
	str := sc.Text()
	res, err := strconv.Atoi(str)
	return int(res), err
}

func validTicket(num int) bool {
	return 1e5 <= num && num <= 1e6
}
