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
	scan := func(q string) int {
		if err != nil {
			return 0
		}
		sc := bufio.NewScanner(os.Stdin)
		fmt.Print(q)
		sc.Scan()
		data := sc.Text()
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
	call := 0
	one := 0
	two := 1
	return func() int {
		switch call {
		case 0:
			call++
			return one
		case 1:
			call++
			return two
		default:
			tmp := one + two
			one = two
			two = tmp
			return two
		}
	}
}
