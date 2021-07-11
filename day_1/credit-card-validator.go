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

	if isValidCard(data) {
		fmt.Println("valid")
		return
	}

	fmt.Println("not valid")
}

func isValidCard(card string) bool {
	splitted := strings.Split(card, " ")
	if len(splitted) != 4 {
		return false
	}

	// except last 4
	array := splitted[0:3]

	for _, elem := range array {
		if len(elem) != 4 {
			return false
		}

		num, err := strconv.Atoi(elem)
		if err != nil {
			return false
		}

		if num < 0 {
			return false
		}
	}

	return true
}
