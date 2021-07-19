package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const helloMsg = `mode 1: count the number of occurrences of a string 
mode 12: replace the string with a new string 
program takes such arguments on the start:
1. <file path> <string to count>
2. <file path> <string to find> <string to replace>`

func main() {
	fmt.Println(helloMsg, "\n")

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
		fmt.Println(fmt.Errorf("File read: %w", err))
	}
	defer file.Close()

	if p == 2 {
		c := 0
		sc := bufio.NewScanner(file)

		for sc.Scan() {
			txt := sc.Text()
			c += strings.Count(txt, substr)
		}

		if err := sc.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(fmt.Sprintf("appeared %d times", c))
		return
	}

	newSubstr := params[2]

	textBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	txt := string(textBytes)
	txt = strings.ReplaceAll(txt, substr, newSubstr)
	ioutil.WriteFile(fileN, []byte(txt), 0644)

	fmt.Println(fmt.Sprintf("replaced %s with %s", substr, newSubstr))
}
