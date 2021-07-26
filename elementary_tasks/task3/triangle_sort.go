package task3

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type Triangle struct {
	a, b, c, Square float64
	name            string
}

func (it *Triangle) Exists() bool {
	return (it.a+it.b > it.c) && (it.a+it.c > it.b) && (it.b+it.c > it.a)
}

func NewTriangle(a, b, c float64, name string) *Triangle {
	p := (a + b + c) / 2
	s := math.Sqrt(p * (p - a) * (p - b) * (p - c))
	return &Triangle{a, b, c, s, name}
}

func SortTrinalgesDescSq(trs []Triangle) {
	sort.Slice(trs, func(i, j int) bool {
		return trs[i].Square > trs[j].Square
	})
}

func BuildTriangle(params []string) (*Triangle, bool) {
	ok := true
	parseLength := func(str string) float64 {
		if !ok {
			return 0
		}
		l, err := strconv.ParseFloat(str, 64)
		ok = (err == nil) && (l > 0)

		return float64(l)
	}

	name := params[0]
	a := parseLength(params[1])
	b := parseLength(params[2])
	c := parseLength(params[3])

	return NewTriangle(a, b, c, name), ok
}

func ParseParams(str string) ([]string, bool) {
	params := strings.Split(str, ",")
	if len(params) != 4 {
		return nil, false
	}

	return TrimSpaces(params), true
}

func TrimSpaces(params []string) []string {
	for i, elem := range params {
		params[i] = strings.TrimSpace(elem)
	}
	return params
}
