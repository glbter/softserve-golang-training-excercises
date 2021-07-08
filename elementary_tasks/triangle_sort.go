package main

import (
	"bufio"
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

		for i, elem := range params {
			params[i] = strings.TrimSpace(elem)
		}
		name = params[0]
		a, err := strconv.ParseFloat((params[1]), 32)
		if err != nil {
			printRules(scanner)
			continue
		}

		b, err := strconv.ParseFloat((params[2]), 32)
		if err != nil {
			printRules(scanner)
			continue
		}

		c, err := strconv.ParseFloat((params[3]), 32)
		if err != nil {
			printRules(scanner)
			continue
		}

		if a < 0 || b < 0 || c < 0 {
			printRules(scanner)
			continue
		}

		if a+b < c || a+c < b || b+c < a {
			fmt.Println("triangle does not exist")
			askContinue(scanner)
			continue
		}

		triangles = append(triangles, NewTriangle(a, b, c, name))
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
