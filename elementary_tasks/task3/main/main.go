package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/console/print"
	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/console/scan"
	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/task3"
)

const helloMsg = `program takes triangles as input by convention:
<name>, <the length of a side a>, <the length of a side b>, <the length of a side c>
to finish input print "n"
program will output the list of trianles sorted by their square in descending order `
const rule = "you should type a name and three positive numbers"

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
		l, e := scan.ScanPositiveFloat(sc, "")
		err = e
		return float64(l)
	}

	for run {
		err = nil
		data := scan.ScanString(sc, "")
		if strings.TrimSpace(data) == "n" {
			break
		}

		params := strings.Split(data, ",")
		fmt.Println(params)
		if len(params) != 4 {
			fmt.Println("wrong format")
			run = print.PrintInstruction(sc, rule)
			continue
		}
		// todo: fix a bug goto like
		fmt.Println()
		params = task3.TrimSpaces(params)
		fmt.Println(params)
		name := params[0]
		a := parseLength(params[1])
		b := parseLength(params[2])
		c := parseLength(params[3])
		// if err != nil {
		// 	run = print.PrintInstruction(sc, rule)
		// 	continue
		// }

		tr := task3.NewTriangle(a, b, c, name)
		fmt.Println(tr)
		if !tr.Exists() {
			fmt.Println("triangle does not exist")
			run = scan.AskContinue(sc)
			continue
		}

		trs = append(trs, *tr)
	}

	sort.Slice(trs, func(i, j int) bool {
		return trs[i].Square > trs[j].Square
	})

	task3.PrintTriangles(trs)
}
