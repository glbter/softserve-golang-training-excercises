package task4

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

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

type ChangedString struct {
	Old string
	New string
}
