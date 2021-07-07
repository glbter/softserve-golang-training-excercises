package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := scanner.Text()

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

	for i := 1; i*i < n; i++ {
		fmt.Println(i)
	}
}

func printRules() {
	fmt.Println("you should type a positive integer")
}
