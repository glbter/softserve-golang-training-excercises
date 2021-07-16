package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const helloMsg = `you can check whether you can put one envelop into another,
every envelop is defined by height and width,
after one comparison you can continue by typing "y" or "yes."
`

var run bool

func main() {
	fmt.Println(helloMsg)

	sc := bufio.NewScanner(os.Stdin)
	run = true

	var Err error
	consoleScan := func(txt string) float32 {
		if Err != nil {
			return 0
		}
		num, err := scanFloat(sc, txt)
		Err = err
		if num < 0 {
			Err = errors.New("negative lenght")
		}
		return num
	}

	for run {
		fmt.Println("First envelop")
		Err = nil

		env1 := Envelop{consoleScan("height"), consoleScan("width")}

		if Err == nil {
			fmt.Println("Second envelop")
		}

		env2 := Envelop{consoleScan("height"), consoleScan("width")}

		if Err != nil {
			printInstruction(sc)
			continue
		}

		env1 = *env1.makeHeightBigger()
		env2 = *env2.makeHeightBigger()

		if env1.isBiggerThan(&env2) {
			fmt.Println("you can put the first envelope into the second")
		} else if env2.isBiggerThan(&env1) {
			fmt.Println("you can put the second envelope into the first")
		} else {
			fmt.Println("you can't put one envlope into another")
		}

		run = askContinue(sc)
	}
	fmt.Println("shutting down...")
}

func (it *Envelop) isBiggerThan(o *Envelop) bool {
	return it.height > o.height && it.width > o.width
}

type Envelop struct {
	height, width float32
}

func (it *Envelop) makeHeightBigger() *Envelop {
	if it.height < it.width {
		it.height, it.width = it.width, it.height
	}

	return it
}

func scanFloat(sc *bufio.Scanner, txt string) (float32, error) {
	fmt.Print(txt, ":  ")
	sc.Scan()
	data := sc.Text()
	res, err := strconv.ParseFloat((data), 32)
	return float32(res), err
}

func askContinue(sc *bufio.Scanner) bool {
	fmt.Println("Continue? ")
	sc.Scan()
	data := sc.Text()
	doContinue := strings.ToLower(strings.TrimSpace(data))

	return doContinue == "yes" || doContinue == "y"
}

func printInstruction(sc *bufio.Scanner) {
	fmt.Println("you should type in envelope parameters one by one. They should be positive numbers. They are width and size")
	run = askContinue(sc)
}
