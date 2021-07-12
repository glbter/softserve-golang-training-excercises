package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var isRunning bool

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	isRunning = true

	var blockError error
	consoleScan := func(text string) float32 {
		if blockError != nil {
			return 0
		}
		num, err := scanFloat(scanner, text)
		blockError = err
		if num < 0 {
			blockError = errors.New("negative lenght")
		}
		return num
	}

	for isRunning {
		fmt.Println("First envelop")
		blockError = nil

		env1 := Envelop{consoleScan("height"), consoleScan("width")}

		if blockError == nil {
			fmt.Println("Second envelop")
		}

		env2 := Envelop{consoleScan("height"), consoleScan("width")}

		if blockError != nil {
			printInstruction(scanner)
			continue
		}

		env1 = env1.makeHightBigger()
		env2 = env2.makeHightBigger()

		fmt.Println(env1, "\n", env2)

		if env1.isBiggerThat(env2) {
			fmt.Println("you can put the first envelope into the second")
		} else if env2.isBiggerThat(env1) {
			fmt.Println("you can put the second envelope into the first")
		} else {
			fmt.Println("you can't put one envlope into another")
		}

		isRunning = askContinue(scanner)
	}
	fmt.Println("shutting down...")
}

func (it Envelop) isBiggerThat(o Envelop) bool {
	return it.height > o.height && it.width > o.width
}

type Envelop struct {
	height, width float32
}

func (it Envelop) makeHightBigger() Envelop {
	it.height, it.width = swapFirstBigger(it.height, it.width)
	return it
}

func swapFirstBigger(a, b float32) (float32, float32) {
	if b > a {
		return b, a
	}
	return a, b
}

func scanFloat(scanner *bufio.Scanner, text string) (float32, error) {
	fmt.Print(text, ":  ")
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
	fmt.Println("you should type in envelope parameters one by one. They should be positive numbers. They are width and size")
	isRunning = askContinue(scanner)
}
