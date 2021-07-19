package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/console/print"
	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/console/scan"
	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/task7"
)

const helloMsg = `program takes a positive integer as input
outputs all natural numbers which squares are less than input number`

func main() {
	fmt.Println(fmt.Sprint(helloMsg, "\n"))

	sc := bufio.NewScanner(os.Stdin)

	n, err := scan.ScanPositiveInt(sc, "")
	if err != nil {
		printRules()
		return
	}

	nums := task7.NaturalSquaresLessThan(n)
	print.PrintNums(nums)
}

func printRules() {
	fmt.Println("you should type a positive integer")
}
