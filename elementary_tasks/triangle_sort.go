package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var run bool

func main() {
	sc := bufio.NewScanner(os.Stdin)

	trs := make([]Triangle, 1)
	var name string
	run = true
	fmt.Println("type 'n' to stop")

	var err error
	parseLength := func(str string) float64 {
		if err != nil {
			return 0
		}
		l, e := strconv.ParseFloat(str, 32)
		err = e
		if l < 0 {
			err = errors.New("not positibe number")
		}
		return l
	}

	for run {
		sc.Scan()
		data := sc.Text()
		if strings.TrimSpace(data) == "n" {
			run = false
			continue
		}

		params := strings.Split(data, ",")
		if len(params) != 4 {
			fmt.Println("wrong format")
			printRules(sc)
			continue
		}

		params = trimSpaces(params)

		name = params[0]
		a := parseLength(params[1])
		b := parseLength(params[2])
		c := parseLength(params[2])

		tr := NewTriangle(a, b, c, name)

		if !tr.Exists() {
			fmt.Println("triangle does not exist")
			askContinue(sc)
			continue
		}

		trs = append(trs, tr)
	}

	sort.Slice(trs, func(i, j int) bool {
		return trs[i].square > trs[j].square
	})

	printTriangles(trs)
}

type Triangle struct {
	a, b, c, square float64
	name            string
}

func (it *Triangle) Exists() bool {
	return it.a+it.b < it.c || it.a+it.c < it.b || it.b+it.c < it.a
}

func trimSpaces(params []string) []string {
	for i, elem := range params {
		params[i] = strings.TrimSpace(elem)
	}
	return params
}

func NewTriangle(a, b, c float64, name string) Triangle {
	p := (a + b + c) / 2
	s := math.Sqrt(p * (p - a) * (p - b) * (p - c))
	return Triangle{a, b, c, s, name}
}

func printRules(sc *bufio.Scanner) {
	fmt.Println("you should type a name and three positive numbers")
	run = askContinue(sc)
}

func printTriangles(trs []Triangle) {
	fmt.Println("============= Triangles list: ===============")
	for _, tr := range trs {
		if tr.square == 0 {
			break
		}
		fmt.Println(tr.name, " ", tr.square)
	}
}

func askContinue(sc *bufio.Scanner) bool {
	fmt.Println("Continue? ")
	sc.Scan()
	data := sc.Text()
	cont := strings.ToLower(strings.TrimSpace(data))

	return cont == "yes" || cont == "y"
}
