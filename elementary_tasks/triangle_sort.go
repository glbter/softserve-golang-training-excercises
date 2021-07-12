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

var isRunning bool

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	triangles := make([]Triangle, 1)
	var name string
	isRunning = true
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

	for isRunning {
		scanner.Scan()
		data := scanner.Text()
		if strings.TrimSpace(data) == "n" {
			isRunning = false
			continue
		}

		params := strings.Split(data, ",")
		if len(params) != 4 {
			fmt.Println("wrong format")
			printRules(scanner)
			continue
		}

		params = trimSpaces(params)

		name = params[0]
		a := parseLength(params[1])
		b := parseLength(params[2])
		c := parseLength(params[2])

		triangle := NewTriangle(a, b, c, name)

		if !triangle.Exists() {
			fmt.Println("triangle does not exist")
			askContinue(scanner)
			continue
		}

		triangles = append(triangles, triangle)
	}

	sort.Slice(triangles, func(i, j int) bool {
		return triangles[i].square > triangles[j].square
	})

	printTriangles(triangles)
}

type Triangle struct {
	a, b, c, square float64
	name            string
}

func (it Triangle) Exists() bool {
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

func printRules(scanner *bufio.Scanner) {
	fmt.Println("you should type a name and three positive numbers")
	isRunning = askContinue(scanner)
}

func printTriangles(triangles []Triangle) {
	fmt.Println("============= Triangles list: ===============")
	for _, triangle := range triangles {
		if triangle.square == 0 {
			break
		}
		fmt.Println(triangle.name, " ", triangle.square)
	}
}

func askContinue(scanner *bufio.Scanner) bool {
	fmt.Println("Continue? ")
	scanner.Scan()
	data := scanner.Text()
	doContinue := strings.ToLower(strings.TrimSpace(data))

	return doContinue == "yes" || doContinue == "y"
}
