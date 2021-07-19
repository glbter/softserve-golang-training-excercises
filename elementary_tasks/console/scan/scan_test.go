package scan

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanPositiveInt(t *testing.T) {
	q := "int?"
	tts := []struct {
		q    string
		sc   *bufio.Scanner
		want int
		err  bool
	}{
		{q, bufio.NewScanner(bytes.NewBufferString("1")), 1, false},
		{q, bufio.NewScanner(bytes.NewBufferString("-1")), 0, true},
		{q, bufio.NewScanner(bytes.NewBufferString("0")), 0, true},
		{q, bufio.NewScanner(bytes.NewBufferString("1.5")), 0, true},
	}
	for _, tt := range tts {
		got, err := ScanPositiveInt(tt.sc, tt.q)
		assert.Equal(t, err != nil, tt.err)
		if err == nil {
			assert.Equal(t, tt.want, got)
		}
	}
}

func TestAskContinue(t *testing.T) {
	tts := []struct {
		sc   *bufio.Scanner
		want bool
	}{
		{bufio.NewScanner(bytes.NewBufferString("y")), true},
		{bufio.NewScanner(bytes.NewBufferString("yes")), true},
		{bufio.NewScanner(bytes.NewBufferString("Y")), true},
		{bufio.NewScanner(bytes.NewBufferString("YeS")), true},
		{bufio.NewScanner(bytes.NewBufferString("no")), false},
		{bufio.NewScanner(bytes.NewBufferString("yeah")), false},
		{bufio.NewScanner(bytes.NewBufferString("n")), false},
	}
	for _, tt := range tts {
		assert.Equal(t, tt.want, AskContinue(tt.sc))
	}
}

func TestScanPositiveFloat(t *testing.T) {
	tts := []struct {
		sc   *bufio.Scanner
		want float32
		err  bool
	}{
		{bufio.NewScanner(bytes.NewBufferString("a")), 0, true},
		{bufio.NewScanner(bytes.NewBufferString("1")), 1, false},
		{bufio.NewScanner(bytes.NewBufferString("-1")), -1, true},
		{bufio.NewScanner(bytes.NewBufferString("-1.5")), -1.5, true},
		{bufio.NewScanner(bytes.NewBufferString("1.5")), 1.5, false},
	}
	for _, tt := range tts {
		got, err := ScanPositiveFloat(tt.sc, "?")
		assert.Equal(t, tt.err, err != nil)
		if err != nil {
			assert.Equal(t, tt.want, got)
		}
	}
}
