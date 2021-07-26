package task4

import (
	"bufio"
	"errors"
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
	r := strings.NewReader(str)
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	params := make([]string, 0, 3)
	for sc.Scan() {
		params = append(params, sc.Text())
	}

	p := len(params)
	if p != 2 && p != 3 {
		return nil, ErrWrongInputForm
	}

	return params, nil
}
