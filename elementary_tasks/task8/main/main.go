package main

import (
	"fmt"
	"os"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/console/print"
	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/task8"
)

const helloMsg = `you can type a min and max number to create an interval of fibonacci numbers`

func main() {
	fmt.Println(fmt.Sprint(helloMsg, "\n"))

	mm, err := task8.GetMinMax(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := task8.FibInRange(mm)
	print.PrintNums(r)
}
