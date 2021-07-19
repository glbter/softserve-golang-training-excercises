package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/task1"
)

const helloMsg = `you can type height and width of your chessboard
and a symbol to be used in it
`

func main() {
	fmt.Println(helloMsg)

	var err error
	sc := bufio.NewScanner(os.Stdin)
	consoleScan := func(txt string) int {
		if err != nil {
			return 0
		}
		num, e := task1.ScanPositiveInt(sc, txt)
		err = e
		return num
	}

	i := consoleScan("input height: ")
	w := consoleScan("input width: ")
	if err != nil {
		fmt.Println("should be positive integers")
		return
	}

	s := task1.ScanString(sc, "symbol: ")
	row := task1.CreateRow(w)
	chs := task1.MakeChessboard(
		task1.FormatOddRow(row, s),
		task1.FormatEvenRow(row, s),
		i)
	fmt.Println(chs)
}
