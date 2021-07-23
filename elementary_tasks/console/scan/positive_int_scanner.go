package scan

import (
	"bufio"
)

type PositiveIntScanner struct {
	sc  *bufio.Scanner
	err error
}

func NewPositiveIntScanner(sc *bufio.Scanner) *PositiveIntScanner {
	return &PositiveIntScanner{sc: sc}
}

func (pis *PositiveIntScanner) Scan(q string) int {
	if pis.err != nil {
		return 0
	}

	r, err := ScanPositiveInt(pis.sc, q)
	pis.err = err
	return r
}

func (pis *PositiveIntScanner) Err() error {
	return pis.err
}
