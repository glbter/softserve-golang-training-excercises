package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := scanner.Text()

	array := strings.Split(data, ",")

	amount := 0
	for _, elem := range array {
		num, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Println("not a number")
			return
		}

		if num > 0 && num%2 == 0 {
			amount++
		}
	}

	fmt.Println(amount)
}
