package task8

import (
	"bufio"
	"io"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/console/scan"
)

// open range
func FibInRange(mm *MinMax) []int {
	min, max := mm.Min, mm.Max
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

type MinMax struct {
	Min int
	Max int
}

func GetMinMax(r io.Reader) (*MinMax, error) {
	sc := bufio.NewScanner(r)
	pis := scan.NewPositiveIntScanner(sc)

	min := pis.Scan("greater than:  ")
	max := pis.Scan("less than:  ")
	if err := pis.Err(); err != nil {
		return nil, err
	}

	if max < min {
		min, max = max, min
	}

	return &MinMax{min, max}, nil
}
