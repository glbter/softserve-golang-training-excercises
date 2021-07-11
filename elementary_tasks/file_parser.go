package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	consoleScanner := bufio.NewScanner(os.Stdin)
	consoleScanner.Scan()
	data := consoleScanner.Text()
	params := strings.Split(strings.TrimSpace(data), " ")

	numParams := len(params)
	if numParams != 2 && numParams != 3 {
		fmt.Println("wrong format: should be 2 <path> <str to count> or <path> <str to find> <str to replace>")
		return
	}

	fileName := params[0]
	substr := params[1]

	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(fmt.Errorf("File read: %w", err))
	}
	defer file.Close()

	if numParams == 2 {
		count := 0
		fileScanner := bufio.NewScanner(file)

		for fileScanner.Scan() {
			text := fileScanner.Text()
			count += strings.Count(text, substr)
		}

		if err := fileScanner.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(fmt.Sprintf("appeared %d times", count))
		return
	}

	newSubstr := params[2]
	// fileWriter := bufio.NewReadWriter(bufio.NewReader(file), bufio.NewWriter(file))
	textBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	text := string(textBytes)
	text = strings.ReplaceAll(text, substr, newSubstr)
	ioutil.WriteFile(fileName, []byte(text), 0644)
	// for {
	// 	line, err := fileWriter.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 		return
	// 	}

	// 	if strings.Contains(line, substr) {
	// 		line = strings.ReplaceAll(line, substr, newSubstr)
	// 	}
	// 	// _, err = fileWriter.WriteString(line)
	// 	// if err != nil {
	// 	// 	log.Fatalln(err)
	// 	// }

	// 	ioutil.WriteFile(fileName, []byte(line), 0644)
	// 	// fileWriter.Flush()
	// }

	fmt.Println(fmt.Sprintf("replaced %s with %s", substr, newSubstr))
}
