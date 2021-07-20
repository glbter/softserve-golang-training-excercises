package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

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
	data := sc.Text()
	params := strings.Split(strings.TrimSpace(data), " ")

	p := len(params)
	if p != 2 && p != 3 {
		fmt.Println("wrong format: should be 2 <path> <str to count> or <path> <str to find> <str to replace>")
		return
	}

	for i, elem := range params {
		params[i] = strings.TrimSpace(elem)
	}

	fileN := params[0]
	substr := params[1]

	file, err := os.OpenFile(fileN, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(fmt.Errorf("file read: %w", err))
	}
	defer file.Close()

	var fileR io.Reader = (*os.File)(file)
	if p == 2 {
		c := task4.CountOccurancesInFile(&fileR, substr)
		fmt.Printf("appeared %d times \n", c)
		return
	}

	newSubstr := params[2]

	task4.ChangeOccurancesInFile(&fileR, fileN, &task4.ChangedString{substr, newSubstr})
	fmt.Printf("replaced %s with %s \n", substr, newSubstr)
}
