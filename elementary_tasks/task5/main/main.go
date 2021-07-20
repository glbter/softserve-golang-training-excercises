package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/console/scan"
	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/task5"
)

const helloMsg = `program takes a positive integer as input and returns it word representation`

func main() {
	fmt.Println(fmt.Sprint(helloMsg, "\n"))

	sc := bufio.NewScanner(os.Stdin)
	num, err := scan.ScanPositiveInt(sc, "")
	if err != nil && err != scan.ErrNotPosInt {
		fmt.Println("you should type not negative integer")
		return
	}

	fmt.Println(task5.ConvertNumToString(num))
}
