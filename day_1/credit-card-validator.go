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

	splitted := strings.Split(data, " ")
	// except last 4
	array := splitted[0:3]

	for _, elem := range array {
		_, err := strconv.Atoi(elem)
		if err != nil {
			fmt.Println("not valid")
			return
		}

	}

	fmt.Println("valid")
}
