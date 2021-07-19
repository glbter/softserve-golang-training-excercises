package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/task8"
)

const helloMsg = `you can type a min and max number to create an interval of fibonacci numbers`

func main() {
	fmt.Println(helloMsg, "\n")

	var err error
	scan := func(q string) int {
		if err != nil {
			return 0
		}
		sc := bufio.NewScanner(os.Stdin)
		fmt.Print(q)
		sc.Scan()
		data := sc.Text()
		num, e := strconv.Atoi(strings.TrimSpace(data))
		err = e
		if num < 0 {
			err = errors.New("negative number found")
		}
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
	task8.PrintRange(r)
}
