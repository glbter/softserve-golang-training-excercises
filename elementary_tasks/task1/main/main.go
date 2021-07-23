package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/console/scan"
	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/task1"
)

const helloMsg = `you can type height and width of your chessboard
and a symbol to be used in it
`

func main() {
	fmt.Println(helloMsg)

	sc := bufio.NewScanner(os.Stdin)
	ss := scan.NewPositiveIntScanner(sc)
	i := ss.Scan("input height: ")
	w := ss.Scan("input width: ")

	if ss.Err() != nil {
		fmt.Println("should be positive integers")
		return
	}

	s, _ := scan.ScanString(sc, "symbol: ")
	row := task1.CreateRow(w)
	chs := task1.MakeChessboard(
		task1.FormatOddRow(row, s),
		task1.FormatEvenRow(row, s),
		i)
	fmt.Println(chs)
}
