package task1

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRow(t *testing.T) {
	tts := []struct {
		w    int
		want int
	}{
		{6, 2730}, //101010101010
		{5, 682},  //1010101010
	}
	for _, tt := range tts {
		assert.Equal(t, tt.want, CreateRow(tt.w))
	}
}

func TestFormat(t *testing.T) {
	tts := []struct {
		str  string
		sym  string
		want string
	}{
		{"101010101010", "*", "* * * * * * "},
		{"00000", "a", "     "},
		{"11111", "a", "aaaaa"},
	}
	for _, tt := range tts {
		assert.Equal(t, tt.want, Format(tt.str, tt.sym))
	}
}

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

func TestFormatEvenRow(t *testing.T) {
	tts := []struct {
		in   int
		want string
	}{
		{2730, " 1 1 1 1 1 1"},
		{682, " 1 1 1 1 1"},
		{511, " 11111111"},
	}
	for _, tt := range tts {
		got := FormatEvenRow(tt.in, "1")
		assert.Equal(t, tt.want, got)
	}
}

func TestFormatOddRow(t *testing.T) {
	tts := []struct {
		in   int
		want string
	}{
		{2730, "1 1 1 1 1 1 "},
		{682, "1 1 1 1 1 "},
		{511, "111111111"},
	}
	for _, tt := range tts {
		got := FormatOddRow(tt.in, "1")
		assert.Equal(t, tt.want, got)
	}
}

func TestMakeChessboard(t *testing.T) {
	tts := []struct {
		odd  string
		even string
		h    int
		want string
	}{
		{"lalalala", "olololol", 4, "lalalala\nolololol\nlalalala\nolololol"},
		{"101010", "111111", 3, "101010\n111111\n101010"},
	}
	for _, tt := range tts {
		got := MakeChessboard(tt.odd, tt.even, tt.h)
		assert.Equal(t, tt.want, got)
	}
}
