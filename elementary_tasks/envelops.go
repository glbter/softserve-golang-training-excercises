package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var isRunning bool

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	isRunning = true

	for isRunning {
		fmt.Println("First envelop")

		env1_h, err := scanFloat(scanner, "height")
		if err != nil {
			printInstruction(scanner)
			continue
		}

		env1_w, err := scanFloat(scanner, "width")
		if err != nil {
			printInstruction(scanner)
			continue
		}

		fmt.Println("Second envelop")
		env2_h, err := scanFloat(scanner, "height")
		if err != nil {
			printInstruction(scanner)
			continue
		}

		env2_w, err := scanFloat(scanner, "width")
		if err != nil {
			printInstruction(scanner)
			continue
		}

		if env1_h > env2_h && env1_w > env2_w {
			fmt.Println("you can put the first envelope into the second")
		} else if env1_h < env2_h && env1_w < env2_w {
			fmt.Println("you can put the second envelope into the first")
		} else {
			fmt.Println("you can't put one envlope into another")
		}

		isRunning = askContinue(scanner)
	}
	fmt.Println("shutting down...")
}

func scanFloat(scanner *bufio.Scanner, text string) (float32, error) {
	fmt.Println(text)
	scanner.Scan()
	data := scanner.Text()
	res, err := strconv.ParseFloat((data), 32)
	return float32(res), err
}

func askContinue(scanner *bufio.Scanner) bool {
	fmt.Println("Continue? ")
	scanner.Scan()
	data := scanner.Text()
	doContinue := strings.ToLower(strings.TrimSpace(data))

	return doContinue == "yes" || doContinue == "y"
}

func printInstruction(scanner *bufio.Scanner) {
	fmt.Println("you should type in envelope parameters one by one. They should be numbers. They are width and size")
	isRunning = askContinue(scanner)
}
