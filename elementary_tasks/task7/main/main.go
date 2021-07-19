package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/task7"
)

const helloMsg = `program takes a positive integer as input
outputs all natural numbers which squares are less than input number`

func main() {
	fmt.Println(helloMsg, "\n")

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	data := sc.Text()

	res, err := strconv.Atoi(data)
	if err != nil {
		printRules()
		return
	}

	n := int(res)
	if res < 0 {
		printRules()
		return
	}

	nums := task7.NaturalSquaresLessThan(n)
	task7.PrintNums(nums)
}

func printRules() {
	fmt.Println("you should type a positive integer")
}
