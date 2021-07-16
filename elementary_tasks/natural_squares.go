package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	fmt.Println()
	for i := 1; i*i < n; i++ {
		if i != 1 {
			fmt.Print(", ")
		}
		fmt.Print(i)
	}
	fmt.Println()
}

func printRules() {
	fmt.Println("you should type a positive integer")
}
