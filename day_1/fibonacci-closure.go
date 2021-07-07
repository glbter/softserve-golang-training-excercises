package main

import (
	"fmt"
)

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

func main() {
	nextInt := intFibSeq()

	for i := 0; i < 10; i++ {
		fmt.Println(nextInt())
	}
}
