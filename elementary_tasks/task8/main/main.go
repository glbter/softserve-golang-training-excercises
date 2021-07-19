package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/console/print"
	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/console/scan"
	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/task8"
)

const helloMsg = `you can type a min and max number to create an interval of fibonacci numbers`

func main() {
	fmt.Println(fmt.Sprint(helloMsg, "\n"))

	sc := bufio.NewScanner(os.Stdin)
	var err error
	scan := func(q string) int {
		if err != nil {
			return 0
		}

		num, e := scan.ScanPositiveInt(sc, q)
		err = e
		return num
	}

	min := scan("greater than:  ")
	max := scan("less than:  ")
	if err != nil {
		fmt.Println(err)
	}

	if max < min {
		min, max = max, min
	}

	r := task8.FibInRange(min, max)
	print.PrintNums(r)
}
