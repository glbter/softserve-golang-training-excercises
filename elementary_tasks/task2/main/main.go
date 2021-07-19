package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/console/print"
	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/console/scan"
	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/task2"
)

const helloMsg = `you can check whether you can put one envelop into another,
every envelop is defined by height and width,
after one comparison you can continue by typing "y" or "yes."
`
const instruction = "you should type in envelope parameters one by one. They should be positive numbers. They are width and size"

func main() {
	fmt.Println(helloMsg)

	sc := bufio.NewScanner(os.Stdin)
	run := true

	var Err error
	consoleScan := func(txt string) float32 {
		if Err != nil {
			return 0
		}
		num, err := scan.ScanPositiveFloat(sc, txt+": ")
		Err = err
		return num
	}

	for run {
		fmt.Println("First envelop")

		env1 := task2.NewEnvelop(consoleScan("height"), consoleScan("width"))
		if Err == nil {
			fmt.Println("Second envelop")
		}

		env2 := task2.NewEnvelop(consoleScan("height"), consoleScan("width"))
		if Err != nil {
			run = print.PrintInstruction(sc, instruction)
			continue
		}

		fmt.Println(task2.EnvelopComparisonStr(env1, env2))
		run = scan.AskContinue(sc)
	}
	fmt.Println("shutting down...")
}
