package scan

import (
	"bufio"
)

type PositiveFloatScanner struct {
	sc  *bufio.Scanner
	err error
}

func NewPositiveFloatScanner(sc *bufio.Scanner) *PositiveFloatScanner {
	return &PositiveFloatScanner{sc: sc}
}

func (pfs *PositiveFloatScanner) Scan(q string) float32 {
	if pfs.err != nil {
		return 0
	}

	r, err := ScanPositiveFloat(pfs.sc, q)
	pfs.err = err
	return r
}

func (pfs *PositiveFloatScanner) Err() error {
	return pfs.err
}
