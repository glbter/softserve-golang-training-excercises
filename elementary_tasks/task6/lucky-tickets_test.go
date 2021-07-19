package task6

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type formulaTest struct {
	in   []int
	want bool
}

func TestValidateEasyFormula(t *testing.T) {
	tts := []formulaTest{
		{[]int{3, 5, 7, 4, 6, 5}, true},
		{[]int{1, 9, 3, 3, 4, 6}, true},
		{[]int{1, 2, 3, 4, 5, 6}, false},
		{[]int{9, 8, 7, 6, 5, 4}, false},
	}
	for _, tt := range tts {
		assert.Equal(t, tt.want, ValidateEasyFormula(tt.in), fmt.Sprintf("arr %v", tt.in))
	}
}

func TestValidateHardFormula(t *testing.T) {
	tts := []formulaTest{
		{[]int{4, 2, 3, 4, 5, 6}, true},
		{[]int{1, 2, 3, 5, 5, 2}, true},
		{[]int{1, 3, 2, 4, 5, 6}, false},
		{[]int{6, 1, 2, 5, 3, 4}, false},
	}
	for _, tt := range tts {
		assert.Equal(t, tt.want, ValidateHardFormula(tt.in))
	}
}

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

// func TestGetAlgorithm(t *testing.T) {
// 	var easy = ValidateEasyFormula
// 	var hard = ValidateHardFormula
// 	tts := []struct {
// 		algo string
// 		want func([]int) bool
// 	}{
// 		{"piter", hard},
// 		{"moskow", easy},
// 		{"Piter", hard},
// 		{"Moskow", easy},
// 		{"PITER", hard},
// 		{"MOSKOW", easy},
// 		{"pit", nil},
// 		{"mosk", nil},
// 	}
// 	for _, tt := range tts {
// 		got, _ := GetAlgorithm(tt.algo)
// 		// assert.Nil(t, err)
// 		// if &got != &tt.want {
// 		// 	t.Errorf("GetAlgorithm(%q), test %d failed", tt.algo, i)
// 		// 	t.Error(&got, &tt.want)
// 		// }
// 		assert.Equal(t, &tt.want, &got)
// 	}
// }
