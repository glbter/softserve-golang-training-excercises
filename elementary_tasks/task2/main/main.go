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
	progCycle(sc)

	fmt.Println("shutting down...")
}

func progCycle(sc *bufio.Scanner) {
	qh, qw := "height: ", "width: "

	run := true
	for run {
		fmt.Println("First envelop")
		pfs := scan.NewPositiveFloatScanner(sc)

		env1 := task2.NewEnvelop(pfs.Scan(qh), pfs.Scan(qw))
		if pfs.Err() == nil {
			fmt.Println("Second envelop")
		}

		env2 := task2.NewEnvelop(pfs.Scan(qh), pfs.Scan(qw))
		if pfs.Err() != nil {
			run = print.PrintInstruction(sc, instruction)
			continue
		}

		fmt.Println(task2.EnvelopComparisonStr(env1, env2))
		run = scan.AskContinue(sc)
	}
}
