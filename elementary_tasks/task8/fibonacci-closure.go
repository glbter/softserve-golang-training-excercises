package task8

import (
	"bufio"
	"io"
	"strconv"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/console/scan"
)

// closed range
func FibInRange(mm *MinMax) []int {
	min, max := mm.Min, mm.Max
	ar := make([]int, 0)

	nextInt := intFibSeq()
	num := nextInt()
	for num <= max {
		if min <= num {
			ar = append(ar, num)
		}
		num = nextInt()
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
		default:
			one, two = two, one+two
			return one
		}
	}
}

type MinMax struct {
	Min, Max int
}

func GetMinMax(r io.Reader) (*MinMax, error) {
	sc := bufio.NewScanner(r)

	minS, _ := scan.ScanString(sc, "greater than:  ")
	maxS, _ := scan.ScanString(sc, "less than:  ")
	min, err := strconv.Atoi(minS)
	if err != nil {
		return nil, err
	}
	max, err := strconv.Atoi(maxS)
	if err != nil {
		return nil, err
	}

	if max < min {
		min, max = max, min
	}

	return &MinMax{min, max}, nil
}
