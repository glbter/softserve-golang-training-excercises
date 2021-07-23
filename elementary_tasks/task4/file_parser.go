package task4

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

var ErrWrongInputForm = errors.New("wrong format: should be 2 <path> <str to count> or <path> <str to find> <str to replace>")

type ChangedString struct {
	Old string
	New string
}

func CountOccurancesInFile(r *io.Reader, substr string) int {
	c := 0
	sc := bufio.NewScanner(*r)

	for sc.Scan() {
		txt := sc.Text()
		c += strings.Count(txt, substr)
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return c
}

func ChangeOccurancesInFile(r *io.Reader, file string, cs *ChangedString) {
	textBytes, err := ioutil.ReadAll(*r)
	if err != nil {
		log.Fatalln(err)
	}

	txt := string(textBytes)
	txt = strings.ReplaceAll(txt, cs.Old, cs.New)
	ioutil.WriteFile(file, []byte(txt), 0644)
}

func GetParams(str string) ([]string, error) {
	params := strings.Split(strings.TrimSpace(str), " ")
	p := len(params)
	if p != 2 && p != 3 {
		return nil, ErrWrongInputForm
	}

	for i, elem := range params {
		params[i] = strings.TrimSpace(elem)
	}

	return params, nil
}

func FileHandleStr(r io.Reader, params []string) string {
	file := params[0]
	substr := params[1]

	if len(params) == 2 {
		c := CountOccurancesInFile(&r, substr)
		return fmt.Sprintf("appeared %d times \n", c)
	}

	newSubstr := params[2]
	ChangeOccurancesInFile(&r, file, &ChangedString{Old: substr, New: newSubstr})
	return fmt.Sprintf("replaced %s with %s \n", substr, newSubstr)
}
