package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/task2"
)

const helloMsg = `you can check whether you can put one envelop into another,
every envelop is defined by height and width,
after one comparison you can continue by typing "y" or "yes."
`

var run bool

func main() {
	fmt.Println(helloMsg)

	sc := bufio.NewScanner(os.Stdin)
	run = true

	var Err error
	consoleScan := func(txt string) float32 {
		if Err != nil {
			return 0
		}
		num, err := task2.ScanPositiveFloat(sc, txt)
		Err = err
		return num
	}

	for run {
		fmt.Println("First envelop")
		Err = nil

		env1 := task2.NewEnvelop(consoleScan("height"), consoleScan("width"))
		if Err == nil {
			fmt.Println("Second envelop")
		}

		env2 := task2.NewEnvelop(consoleScan("height"), consoleScan("width"))
		if Err != nil {
			run = task2.PrintInstruction(sc)
			continue
		}

		fmt.Println(task2.EnvelopComparisonStr(env1, env2))
		run = task2.AskContinue(sc)
	}
	fmt.Println("shutting down...")
}
