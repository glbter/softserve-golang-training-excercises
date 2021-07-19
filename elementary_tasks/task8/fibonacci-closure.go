package task8

import (
	"fmt"
)

func PrintRange(r []int) {
	for i, elem := range r {
		if i != 0 {
			fmt.Print(", ")
		}
		fmt.Print(elem)
	}
	fmt.Println()
}

// open range
func FibInRange(min, max int) []int {
	ar := make([]int, 0)
	num := 0
	nextInt := intFibSeq()
	for num < max {
		num = nextInt()
		if min < num && num < max {
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
