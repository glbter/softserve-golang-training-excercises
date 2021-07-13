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
	scan := func(question string) int {
		if err != nil {
			return 0
		}
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(question)
		scanner.Scan()
		data := scanner.Text()
		num, e := strconv.Atoi(strings.TrimSpace(data))
		err = e
		if num < 0 {
			err = errors.New("negative number found")
		}
		return num
	}

	min := scan("greater than:  ")
	max := scan("less than:  ")
	if err != nil {
		fmt.Println(err)
	}

	if max < min {
		min, max = max, min
	}

	r := fibInRange(min, max)
	printRange(r)
}

func printRange(r []int) {
	for _, elem := range r {
		fmt.Print(elem, ", ")
	}
	fmt.Println()
}

func fibInRange(min, max int) []int {
	ar := make([]int, 0)
	num := 0
	nextInt := intFibSeq()
	for num < max {
		num = nextInt()
		if num > max {
			return ar
		}
		if num > min {
			ar = append(ar, num)
		}
	}
	return ar
}

func intFibSeq() func() int {
	call_num := 0
	first := 0
	second := 1
	return func() int {
		switch call_num {
		case 0:
			call_num++
			return first
		case 1:
			call_num++
			return second
		default:
			tmp := first + second
			first = second
			second = tmp
			return second
		}
	}
}
