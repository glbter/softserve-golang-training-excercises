package task2

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const FirstIntoSecond = "you can put the first envelope into the second"
const SecondIntoFirst = "you can put the second envelope into the first"
const CantCompare = "you can't put one envlope into another"

type Envelop struct {
	height, width float32
}

func (it *Envelop) IsBiggerThan(o *Envelop) bool {
	return it.height > o.height && it.width > o.width
}

func (it *Envelop) MakeHeightBigger() *Envelop {
	if it.height < it.width {
		it.height, it.width = it.width, it.height
	}

	return it
}

func NewEnvelop(h, w float32) *Envelop {
	e := Envelop{h, w}
	return e.MakeHeightBigger()
}

func EnvelopComparisonStr(env1, env2 *Envelop) string {
	if env1.IsBiggerThan(env2) {
		return SecondIntoFirst
	} else if env2.IsBiggerThan(env1) {
		return FirstIntoSecond
	}
	return CantCompare
}

func ScanPositiveFloat(sc *bufio.Scanner, txt string) (float32, error) {
	fmt.Print(txt, ":  ")
	sc.Scan()
	data := sc.Text()
	res, err := strconv.ParseFloat((data), 32)
	if res < 0 {
		err = errors.New("negative lenght")
	}
	return float32(res), err
}

func AskContinue(sc *bufio.Scanner) bool {
	fmt.Println("Continue? ")
	sc.Scan()
	data := sc.Text()
	doContinue := strings.ToLower(strings.TrimSpace(data))

	return doContinue == "yes" || doContinue == "y"
}

func PrintInstruction(sc *bufio.Scanner) bool {
	fmt.Println("you should type in envelope parameters one by one. They should be positive numbers. They are width and size")
	return AskContinue(sc)
}
