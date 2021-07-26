package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/glbter/softserve-golang-training-excercises/elementary_tasks/task4"
)

const helloMsg = `mode 1: count the number of occurrences of a string 
mode 12: replace the string with a new string 
program takes such arguments on the start:
1. <file path> <string to count>
2. <file path> <string to find> <string to replace>`

func main() {
	fmt.Println(fmt.Sprint(helloMsg, "\n"))

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	params, err := task4.GetParams(sc.Text())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fileN := params[0]
	file, err := os.OpenFile(fileN, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(fmt.Errorf("file read: %w", err))
		return
	}
	defer file.Close()

	fmt.Println(fileHandleStr(file, params))
}

func fileHandleStr(r io.Reader, params []string) string {
	file := params[0]
	substr := params[1]

	if len(params) == 2 {
		c := task4.CountOccurancesInFile(&r, substr)
		return fmt.Sprintf("appeared %d times \n", c)
	}

	newSubstr := params[2]
	task4.ChangeOccurancesInFile(&r, file, &task4.ChangedString{Old: substr, New: newSubstr})
	return fmt.Sprintf("replaced %s with %s \n", substr, newSubstr)
}
