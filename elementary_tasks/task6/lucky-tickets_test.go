package task6

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	f "github.com/glbter/softserve-golang-training-excercises/elementary_tasks/task6/formula"
)

func TestValidTicket(t *testing.T) {
	tts := []struct {
		ticket int
		valid  bool
	}{
		{4567, false},
		{12345, false},
		{123456, true},
		{987434, true},
		{1234567, false},
		{012345, false}, // 012345 - is not a six digit number, it's a five digit 12345
	}

	for _, tt := range tts {
		got := ValidTicket(tt.ticket)
		assert.Equal(t, tt.valid, got, "ticket should be six digit number")
	}
}

func TestGetAlgorithm(t *testing.T) {
	tts := []struct {
		algo string
		want FormulaValidator
	}{
		{"piter", &f.HardFormula{}},
		{"moskow", &f.EasyFormula{}},
		{"Piter", &f.HardFormula{}},
		{"Moskow", &f.EasyFormula{}},
		{"PITER", &f.HardFormula{}},
		{"MOSKOW", &f.EasyFormula{}},
		{"pit", nil},
		{"mosk", nil},
	}
	for _, tt := range tts {
		got, _ := GetAlgorithm(tt.algo)
		assert.Equal(t, tt.want, got, fmt.Sprintf("for input: %v", tt.algo))
	}
}

func TestCountTickets(t *testing.T) {
	str := " 123456 \n 357465 \n 193346 \n 987654 \n 423456 \n 123552 \n 132456 \n 612534 "
	// 357465 - easy
	// 193346 - easy
	// 423456 - hard
	// 123552 - hard
	sc1 := bufio.NewScanner(bytes.NewReader([]byte(str)))
	sc2 := bufio.NewScanner(bytes.NewBufferString(str))

	tts := []struct {
		sc   *bufio.Scanner
		alg  FormulaValidator
		want int
	}{
		{sc1, &f.EasyFormula{}, 2},
		{sc2, &f.HardFormula{}, 2},
	}

	for i, tt := range tts {
		got := CountTickets(tt.sc, tt.alg)
		assert.Equal(t, tt.want, got, fmt.Sprintf("testcase num: %d", i))
	}
}
