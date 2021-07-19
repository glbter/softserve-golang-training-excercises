package task7

import (
	"fmt"
)

func NaturalSquaresLessThan(n int) []int {
	r := make([]int, 0)
	for i := 1; i*i < n; i++ {
		r = append(r, i)
	}
	return r
}

func PrintNums(nums []int) {
	fmt.Println()
	for i, num := range nums {
		if i != 0 {
			fmt.Print(", ")
		}
		fmt.Print(num)
	}
	fmt.Println(".")
}
