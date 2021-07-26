package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
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

var (
	errWrongFormat = errors.New("wrong format, should be: <name>, <the length of a side a>, <the length of a side b>, <the length of a side c>")
	errNegativeNum = errors.New("all lengths should pe positive numbers")
)

type triangleError struct {
	tr task3.Triangle
}

func (te *triangleError) Error() string {
	return fmt.Sprintf("such triangle can't exist %+v", te.tr)
}

func main() {
	fmt.Println(fmt.Sprint(helloMsg, "\n"))

	trs := scanTrianglesList(os.Stdin)

	task3.SortTrinalgesDescSq(trs)

	task3.PrintTriangles(trs)
}

func scanTrianglesList(r io.Reader) []task3.Triangle {
	sc := bufio.NewScanner(r)
	trs := make([]task3.Triangle, 0)
	run := true
	fmt.Println("type 'n' to stop")

	for run {
		data, _ := scan.ScanString(sc, "")
		if strings.TrimSpace(data) == "n" {
			break
		}

		tr, err := parseTriangle(data)
		if err == errWrongFormat || err == errNegativeNum {
			fmt.Println(err.Error())
			run = print.PrintInstruction(sc, rule)
			continue
		}
		if err != nil {
			fmt.Println(err.Error())
			run = scan.AskContinue(sc)
			continue
		}

		trs = append(trs, *tr)
	}

	return trs
}

func parseTriangle(str string) (*task3.Triangle, error) {
	params, ok := task3.ParseParams(str)
	if !ok {
		return nil, errWrongFormat
	}

	tr, ok := task3.BuildTriangle(params)
	if !ok {
		return nil, errNegativeNum
	}

	if !tr.Exists() {
		fmt.Println("triangle does not exist")
		return nil, &triangleError{*tr}

	}
	return tr, nil
}
