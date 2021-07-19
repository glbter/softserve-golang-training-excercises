package print

import (
	"bufio"
	"fmt"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/console/scan"
)

func PrintInstruction(sc *bufio.Scanner, str string) bool {
	fmt.Println(str)
	return scan.AskContinue(sc)
}

func PrintNums(nums []int) {
	fmt.Println()
	for i, num := range nums {
		if i != 0 {
			fmt.Print(", ")
		}
		fmt.Print(num)
	}
	fmt.Println(".")
}
