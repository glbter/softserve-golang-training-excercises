package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/task3"
)

const helloMsg = `program takes triangles as input by convention:
<name>, <the length of a side a>, <the length of a side b>, <the length of a side c>
to finish input print "n"
program will output the list of trianles sorted by their square in descending order `

func main() {
	fmt.Println(helloMsg, "\n")

	sc := bufio.NewScanner(os.Stdin)

	trs := make([]task3.Triangle, 1)
	run := true
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
			break
		}

		params := strings.Split(data, ",")
		if len(params) != 4 {
			fmt.Println("wrong format")
			run = task3.PrintRules(sc)
			continue
		}

		params = task3.TrimSpaces(params)

		name := params[0]
		a := parseLength(params[1])
		b := parseLength(params[2])
		c := parseLength(params[3])

		tr := task3.NewTriangle(a, b, c, name)

		if !tr.Exists() {
			fmt.Println("triangle does not exist")
			run = task3.AskContinue(sc)
			continue
		}

		trs = append(trs, *tr)
	}

	sort.Slice(trs, func(i, j int) bool {
		return trs[i].Square > trs[j].Square
	})

	task3.PrintTriangles(trs)
}
