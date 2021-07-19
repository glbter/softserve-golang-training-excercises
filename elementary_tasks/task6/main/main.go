package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/task6"
)

const helloMsg = `program takes a file path as input
in the first line in file should be defined algorithm to count lucky tickets: moskow or piter
than line by line defined six number tickets`

func main() {
	fmt.Println(helloMsg)

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	path := sc.Text()

	file, err := os.Open(strings.TrimSpace(path))
	if err != nil {
		fmt.Println("could not open the file")
		return
	}
	defer file.Close()

	sc = bufio.NewScanner(file)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	algoN := sc.Text()

	algo, err := task6.GetAlgorithm(algoN)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var count int
	for sc.Scan() {
		data, err := task6.Scan(sc)
		if err != nil || !task6.ValidTicket(data) {
			continue
		}

		lucky := algo(task6.SplitToDigits(data))
		if lucky {
			count++
		}
	}

	fmt.Println("--Result--")
	fmt.Println(fmt.Sprint("Amount of lucky tickets: ", count))
}
