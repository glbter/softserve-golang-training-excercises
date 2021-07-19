package task3

import (
	"bufio"
	"fmt"
	"math"
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

func TrimSpaces(params []string) []string {
	for i, elem := range params {
		params[i] = strings.TrimSpace(elem)
	}
	return params
}

func PrintTriangles(trs []Triangle) {
	fmt.Println("============= Triangles list: ===============")
	for _, tr := range trs {
		if tr.Square == 0 {
			break
		}
		fmt.Println(tr.name, " ", tr.Square)
	}
}

func PrintRules(sc *bufio.Scanner) bool {
	fmt.Println("you should type a name and three positive numbers")
	return AskContinue(sc)
}

func AskContinue(sc *bufio.Scanner) bool {
	fmt.Println("Continue? ")
	sc.Scan()
	data := sc.Text()
	cont := strings.ToLower(strings.TrimSpace(data))

	return cont == "yes" || cont == "y"
}
